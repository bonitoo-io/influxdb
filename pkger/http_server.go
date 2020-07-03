package pkger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/influxdata/influxdb/v2"
	pctx "github.com/influxdata/influxdb/v2/context"
	ierrors "github.com/influxdata/influxdb/v2/kit/errors"
	kithttp "github.com/influxdata/influxdb/v2/kit/transport/http"
	"github.com/influxdata/influxdb/v2/pkg/jsonnet"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const RoutePrefix = "/api/v2/packages"

// HTTPServer is a server that manages the packages HTTP transport.
type HTTPServer struct {
	chi.Router
	api    *kithttp.API
	logger *zap.Logger
	svc    SVC
}

// NewHTTPServer constructs a new http server.
func NewHTTPServer(log *zap.Logger, svc SVC) *HTTPServer {
	svr := &HTTPServer{
		api:    kithttp.NewAPI(kithttp.WithLog(log)),
		logger: log,
		svc:    svc,
	}

	exportAllowContentTypes := middleware.AllowContentType("text/yml", "application/x-yaml", "application/json")
	setJSONContentType := middleware.SetHeader("Content-Type", "application/json; charset=utf-8")

	r := chi.NewRouter()
	{
		r.With(exportAllowContentTypes).Post("/", svr.export)
		r.With(setJSONContentType).Post("/apply", svr.apply)

		r.Route("/stacks", func(r chi.Router) {
			r.Post("/", svr.createStack)
			r.Get("/", svr.listStacks)

			r.Route("/{stack_id}", func(r chi.Router) {
				r.Get("/", svr.readStack)
				r.Delete("/", svr.deleteStack)
				r.Patch("/", svr.updateStack)
				r.With(
					exportAllowContentTypes,
					// sunsetting, will not appear as part of the swagger doc and will be dropped
					// in the near term. The sunset header follows the following IETF RFC:
					//	https://tools.ietf.org/html/rfc8594
					middleware.SetHeader("Sunset", "Thurs, 30 July 2020 17:00:00 UTC"),
				).Get("/export", svr.exportStack)
			})
		})
	}

	svr.Router = r
	return svr
}

// Prefix provides the prefix to this route tree.
func (s *HTTPServer) Prefix() string {
	return RoutePrefix
}

type (
	// RespStack is the response body for a stack.
	RespStack struct {
		ID          string              `json:"id"`
		OrgID       string              `json:"orgID"`
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Resources   []RespStackResource `json:"resources"`
		Sources     []string            `json:"sources"`
		URLs        []string            `json:"urls"`
		influxdb.CRUDLog
	}

	// RespStackResource is the response for a stack resource. This type exists
	// to decouple the internal service implementation from the deprecates usage
	// of pkgs in the API. We could add a custom UnmarshalJSON method, but
	// I would rather keep it obvious and explicit with a separate field.
	RespStackResource struct {
		APIVersion   string                   `json:"apiVersion"`
		ID           string                   `json:"resourceID"`
		Kind         Kind                     `json:"kind"`
		MetaName     string                   `json:"templateMetaName"`
		Associations []RespStackResourceAssoc `json:"associations"`

		// PkgName is deprecated moving forward, will support until it is
		// ripped out.
		PkgName *string `json:"pkgName,omitempty"`
	}

	// RespStackResourceAssoc is the response for a stack resource's associations.
	RespStackResourceAssoc struct {
		Kind     Kind   `json:"kind"`
		MetaName string `json:"metaName"`

		//PkgName is to be deprecated moving forward
		PkgName *string `json:"pkgName,omitempty"`
	}
)

// RespListStacks is the HTTP response for a stack list call.
type RespListStacks struct {
	Stacks []RespStack `json:"stacks"`
}

func (s *HTTPServer) listStacks(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	rawOrgID := q.Get("orgID")
	orgID, err := influxdb.IDFromString(rawOrgID)
	if err != nil {
		s.api.Err(w, r, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  fmt.Sprintf("organization id[%q] is invalid", rawOrgID),
			Err:  err,
		})
		return
	}

	if err := r.ParseForm(); err != nil {
		s.api.Err(w, r, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "failed to parse form from encoded url",
			Err:  err,
		})
		return
	}

	filter := ListFilter{
		Names: r.Form["name"],
	}

	for _, idRaw := range r.Form["stackID"] {
		id, err := influxdb.IDFromString(idRaw)
		if err != nil {
			s.api.Err(w, r, &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  fmt.Sprintf("stack ID[%q] provided is invalid", idRaw),
				Err:  err,
			})
			return
		}
		filter.StackIDs = append(filter.StackIDs, *id)
	}

	stacks, err := s.svc.ListStacks(r.Context(), *orgID, filter)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}
	if stacks == nil {
		stacks = []Stack{}
	}

	out := make([]RespStack, 0, len(stacks))
	for _, st := range stacks {
		out = append(out, convertStackToRespStack(st))
	}

	s.api.Respond(w, r, http.StatusOK, RespListStacks{
		Stacks: out,
	})
}

// ReqCreateStack is a request body for a create stack call.
type ReqCreateStack struct {
	OrgID       string   `json:"orgID"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	URLs        []string `json:"urls"`
}

// OK validates the request body is valid.
func (r *ReqCreateStack) OK() error {
	// TODO: provide multiple errors back for failing validation
	if _, err := influxdb.IDFromString(r.OrgID); err != nil {
		return &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  fmt.Sprintf("provided org id[%q] is invalid", r.OrgID),
		}
	}

	for _, u := range r.URLs {
		if _, err := url.Parse(u); err != nil {
			return &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  fmt.Sprintf("provided url[%q] is invalid", u),
			}
		}
	}
	return nil
}

func (r *ReqCreateStack) orgID() influxdb.ID {
	orgID, _ := influxdb.IDFromString(r.OrgID)
	return *orgID
}

func (s *HTTPServer) createStack(w http.ResponseWriter, r *http.Request) {
	var reqBody ReqCreateStack
	if err := s.api.DecodeJSON(r.Body, &reqBody); err != nil {
		s.api.Err(w, r, err)
		return
	}
	defer r.Body.Close()

	auth, err := pctx.GetAuthorizer(r.Context())
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	stack, err := s.svc.InitStack(r.Context(), auth.GetUserID(), Stack{
		OrgID:       reqBody.orgID(),
		Name:        reqBody.Name,
		Description: reqBody.Description,
		URLs:        reqBody.URLs,
	})
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	s.api.Respond(w, r, http.StatusCreated, convertStackToRespStack(stack))
}

func (s *HTTPServer) deleteStack(w http.ResponseWriter, r *http.Request) {
	orgID, err := getRequiredOrgIDFromQuery(r.URL.Query())
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	stackID, err := stackIDFromReq(r)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	auth, err := pctx.GetAuthorizer(r.Context())
	if err != nil {
		s.api.Err(w, r, err)
		return
	}
	userID := auth.GetUserID()

	err = s.svc.DeleteStack(r.Context(), struct{ OrgID, UserID, StackID influxdb.ID }{
		OrgID:   orgID,
		UserID:  userID,
		StackID: stackID,
	})
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	s.api.Respond(w, r, http.StatusNoContent, nil)
}

func (s *HTTPServer) exportStack(w http.ResponseWriter, r *http.Request) {
	stackID, err := stackIDFromReq(r)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	pkg, err := s.svc.Export(r.Context(), ExportWithStackID(stackID))
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	encoding := pkgEncoding(r.Header.Get("Accept"))

	b, err := pkg.Encode(encoding)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	switch encoding {
	case EncodingYAML:
		w.Header().Set("Content-Type", "application/x-yaml")
	default:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	s.api.Write(w, http.StatusOK, b)
}

func (s *HTTPServer) readStack(w http.ResponseWriter, r *http.Request) {
	stackID, err := stackIDFromReq(r)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	stack, err := s.svc.ReadStack(r.Context(), stackID)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	s.api.Respond(w, r, http.StatusOK, convertStackToRespStack(stack))
}

type (
	// ReqUpdateStack is the request body for updating a stack.
	ReqUpdateStack struct {
		Name                *string                  `json:"name"`
		Description         *string                  `json:"description"`
		TemplateURLs        []string                 `json:"templateURLs"`
		AdditionalResources []ReqUpdateStackResource `json:"additionalResources"`

		// Deprecating the urls field and replacing with templateURLs field.
		// This is remaining here for backwards compatibility.
		URLs []string `json:"urls"`
	}

	ReqUpdateStackResource struct {
		ID       string `json:"resourceID"`
		Kind     Kind   `json:"kind"`
		MetaName string `json:"templateMetaName"`
	}
)

func (s *HTTPServer) updateStack(w http.ResponseWriter, r *http.Request) {
	var req ReqUpdateStack
	if err := s.api.DecodeJSON(r.Body, &req); err != nil {
		s.api.Err(w, r, err)
		return
	}

	stackID, err := stackIDFromReq(r)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	update := StackUpdate{
		ID:          stackID,
		Name:        req.Name,
		Description: req.Description,
		URLs:        append(req.TemplateURLs, req.URLs...),
	}
	for _, res := range req.AdditionalResources {
		id, err := influxdb.IDFromString(res.ID)
		if err != nil {
			s.api.Err(w, r, influxErr(influxdb.EInvalid, err, fmt.Sprintf("stack resource id %q", res.ID)))
			return
		}
		update.AdditionalResources = append(update.AdditionalResources, StackAdditionalResource{
			APIVersion: APIVersion,
			ID:         *id,
			Kind:       res.Kind,
			MetaName:   res.MetaName,
		})
	}

	stack, err := s.svc.UpdateStack(r.Context(), update)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	s.api.Respond(w, r, http.StatusOK, convertStackToRespStack(stack))
}

func stackIDFromReq(r *http.Request) (influxdb.ID, error) {
	stackID, err := influxdb.IDFromString(chi.URLParam(r, "stack_id"))
	if err != nil {
		return 0, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "the stack id provided in the path was invalid",
			Err:  err,
		}
	}
	return *stackID, nil
}

func getRequiredOrgIDFromQuery(q url.Values) (influxdb.ID, error) {
	orgIDRaw := q.Get("orgID")
	if orgIDRaw == "" {
		return 0, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "the orgID query param is required",
		}
	}

	orgID, err := influxdb.IDFromString(orgIDRaw)
	if err != nil {
		return 0, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  "the orgID query param was invalid",
			Err:  err,
		}
	}
	return *orgID, nil
}

// ReqExportOrgIDOpt provides options to export resources by organization id.
type ReqExportOrgIDOpt struct {
	OrgID   string `json:"orgID"`
	Filters struct {
		ByLabel        []string `json:"byLabel"`
		ByResourceKind []Kind   `json:"byResourceKind"`
	} `json:"resourceFilters"`
}

// ReqExport is a request body for the export endpoint.
type ReqExport struct {
	StackID   string              `json:"stackID"`
	OrgIDs    []ReqExportOrgIDOpt `json:"orgIDs"`
	Resources []ResourceToClone   `json:"resources"`
}

// OK validates a create request.
func (r *ReqExport) OK() error {
	if len(r.Resources) == 0 && len(r.OrgIDs) == 0 && r.StackID == "" {
		return &influxdb.Error{
			Code: influxdb.EUnprocessableEntity,
			Msg:  "at least 1 resource, 1 org id, or stack id must be provided",
		}
	}

	for _, org := range r.OrgIDs {
		if _, err := influxdb.IDFromString(org.OrgID); err != nil {
			return &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  fmt.Sprintf("provided org id is invalid: %q", org.OrgID),
			}
		}
	}

	if r.StackID != "" {
		_, err := influxdb.IDFromString(r.StackID)
		return err
	}
	return nil
}

// RespExport is a response body for the create pkg endpoint.
type RespExport []Object

func (s *HTTPServer) export(w http.ResponseWriter, r *http.Request) {
	var reqBody ReqExport
	if err := s.api.DecodeJSON(r.Body, &reqBody); err != nil {
		s.api.Err(w, r, err)
		return
	}
	defer r.Body.Close()

	opts := []ExportOptFn{
		ExportWithExistingResources(reqBody.Resources...),
	}
	for _, orgIDStr := range reqBody.OrgIDs {
		orgID, err := influxdb.IDFromString(orgIDStr.OrgID)
		if err != nil {
			continue
		}
		opts = append(opts, ExportWithAllOrgResources(ExportByOrgIDOpt{
			OrgID:         *orgID,
			LabelNames:    orgIDStr.Filters.ByLabel,
			ResourceKinds: orgIDStr.Filters.ByResourceKind,
		}))
	}

	if reqBody.StackID != "" {
		stackID, err := influxdb.IDFromString(reqBody.StackID)
		if err != nil {
			s.api.Err(w, r, &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  fmt.Sprintf("invalid stack ID provided: %q", reqBody.StackID),
			})
			return
		}
		opts = append(opts, ExportWithStackID(*stackID))
	}

	newPkg, err := s.svc.Export(r.Context(), opts...)
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	resp := RespExport(newPkg.Objects)
	if resp == nil {
		resp = []Object{}
	}

	var enc encoder
	switch pkgEncoding(r.Header.Get("Accept")) {
	case EncodingYAML:
		enc = yaml.NewEncoder(w)
		w.Header().Set("Content-Type", "application/x-yaml")
	default:
		enc = newJSONEnc(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	s.encResp(w, r, enc, http.StatusOK, resp)
}

// ReqTemplateRemote provides a package via a remote (i.e. a gist). If content type is not
// provided then the service will do its best to discern the content type of the
// contents.
type ReqTemplateRemote struct {
	URL         string `json:"url" yaml:"url"`
	ContentType string `json:"contentType" yaml:"contentType"`
}

// Encoding returns the encoding type that corresponds to the given content type.
func (p ReqTemplateRemote) Encoding() Encoding {
	return convertEncoding(p.ContentType, p.URL)
}

type ReqRawTemplate struct {
	ContentType string          `json:"contentType" yaml:"contentType"`
	Sources     []string        `json:"sources" yaml:"sources"`
	Pkg         json.RawMessage `json:"contents" yaml:"contents"`
}

func (p ReqRawTemplate) Encoding() Encoding {
	var source string
	if len(p.Sources) > 0 {
		source = p.Sources[0]
	}
	return convertEncoding(p.ContentType, source)
}

// ReqRawAction is a raw action consumers can provide to change the behavior
// of the application of a template.
type ReqRawAction struct {
	Action     string          `json:"action"`
	Properties json.RawMessage `json:"properties"`
}

// ReqApply is the request body for a json or yaml body for the apply pkg endpoint.
type ReqApply struct {
	DryRun  bool                `json:"dryRun" yaml:"dryRun"`
	OrgID   string              `json:"orgID" yaml:"orgID"`
	StackID *string             `json:"stackID" yaml:"stackID"` // optional: non nil value signals stack should be used
	Remotes []ReqTemplateRemote `json:"remotes" yaml:"remotes"`

	// TODO(jsteenb2): pkg references will all be replaced by template references
	// 	these 2 exist alongside the templates for backwards compatibility
	// 	until beta13 rolls out the door. This code should get axed when the next
	//	OSS release goes out.
	RawPkgs []json.RawMessage `json:"packages" yaml:"packages"`
	RawPkg  json.RawMessage   `json:"package" yaml:"package"`

	RawTemplates []ReqRawTemplate `json:"templates" yaml:"templates"`
	RawTemplate  ReqRawTemplate   `json:"template" yaml:"template"`

	EnvRefs map[string]string `json:"envRefs"`
	Secrets map[string]string `json:"secrets"`

	RawActions []ReqRawAction `json:"actions"`
}

// Pkgs returns all pkgs associated with the request.
func (r ReqApply) Pkgs(encoding Encoding) (*Pkg, error) {
	var rawPkgs []*Pkg
	for _, rem := range r.Remotes {
		if rem.URL == "" {
			continue
		}
		pkg, err := Parse(rem.Encoding(), FromHTTPRequest(rem.URL), ValidSkipParseError())
		if err != nil {
			return nil, &influxdb.Error{
				Code: influxdb.EUnprocessableEntity,
				Msg:  fmt.Sprintf("pkg from url[%s] had an issue: %s", rem.URL, err.Error()),
			}
		}
		rawPkgs = append(rawPkgs, pkg)
	}

	for i, rawPkg := range append(r.RawPkgs, r.RawPkg) {
		if rawPkg == nil {
			continue
		}

		pkg, err := Parse(encoding, FromReader(bytes.NewReader(rawPkg)), ValidSkipParseError())
		if err != nil {
			return nil, &influxdb.Error{
				Code: influxdb.EUnprocessableEntity,
				Msg:  fmt.Sprintf("pkg[%d] had an issue: %s", i, err.Error()),
			}
		}
		rawPkgs = append(rawPkgs, pkg)
	}

	for i, rawTmpl := range append(r.RawTemplates, r.RawTemplate) {
		if rawTmpl.Pkg == nil {
			continue
		}
		enc := encoding
		if sourceEncoding := rawTmpl.Encoding(); sourceEncoding != EncodingSource {
			enc = sourceEncoding
		}
		pkg, err := Parse(enc, FromReader(bytes.NewReader(rawTmpl.Pkg), rawTmpl.Sources...), ValidSkipParseError())
		if err != nil {
			sources := formatSources(rawTmpl.Sources)
			return nil, &influxdb.Error{
				Code: influxdb.EUnprocessableEntity,
				Msg:  fmt.Sprintf("pkg[%d] from source(s) %q had an issue: %s", i, sources, err.Error()),
			}
		}
		rawPkgs = append(rawPkgs, pkg)
	}

	return Combine(rawPkgs, ValidWithoutResources(), ValidSkipParseError())
}

type actionType string

// various ActionTypes the transport API speaks
const (
	ActionTypeSkipKind     actionType = "skipKind"
	ActionTypeSkipResource actionType = "skipResource"
)

func (r ReqApply) validActions() (struct {
	SkipKinds     []ActionSkipKind
	SkipResources []ActionSkipResource
}, error) {
	type actions struct {
		SkipKinds     []ActionSkipKind
		SkipResources []ActionSkipResource
	}

	unmarshalErrFn := func(err error, idx int, actionType string) error {
		msg := fmt.Sprintf("failed to unmarshal properties for actions[%d] %q", idx, actionType)
		return ierrors.Wrap(err, msg)
	}

	kindErrFn := func(err error, idx int, actionType string) error {
		msg := fmt.Sprintf("invalid kind for actions[%d] %q", idx, actionType)
		return ierrors.Wrap(err, msg)
	}

	var out actions
	for i, rawAct := range r.RawActions {
		switch a := rawAct.Action; actionType(a) {
		case ActionTypeSkipResource:
			var asr ActionSkipResource
			if err := json.Unmarshal(rawAct.Properties, &asr); err != nil {
				return actions{}, influxErr(influxdb.EInvalid, unmarshalErrFn(err, i, a))
			}
			if err := asr.Kind.OK(); err != nil {
				return actions{}, influxErr(influxdb.EInvalid, kindErrFn(err, i, a))
			}
			out.SkipResources = append(out.SkipResources, asr)
		case ActionTypeSkipKind:
			var ask ActionSkipKind
			if err := json.Unmarshal(rawAct.Properties, &ask); err != nil {
				return actions{}, influxErr(influxdb.EInvalid, unmarshalErrFn(err, i, a))
			}
			if err := ask.Kind.OK(); err != nil {
				return actions{}, influxErr(influxdb.EInvalid, kindErrFn(err, i, a))
			}
			out.SkipKinds = append(out.SkipKinds, ask)
		default:
			msg := fmt.Sprintf(
				"invalid action type %q provided for actions[%d] ; Must be one of [%s]",
				a, i, ActionTypeSkipResource,
			)
			return actions{}, influxErr(influxdb.EInvalid, msg)
		}
	}

	return out, nil
}

// RespApply is the response body for the apply pkg endpoint.
type RespApply struct {
	Sources []string `json:"sources" yaml:"sources"`
	StackID string   `json:"stackID" yaml:"stackID"`
	Diff    Diff     `json:"diff" yaml:"diff"`
	Summary Summary  `json:"summary" yaml:"summary"`

	Errors []ValidationErr `json:"errors,omitempty" yaml:"errors,omitempty"`
}

func (s *HTTPServer) apply(w http.ResponseWriter, r *http.Request) {
	var reqBody ReqApply
	encoding, err := decodeWithEncoding(r, &reqBody)
	if err != nil {
		s.api.Err(w, r, newDecodeErr(encoding.String(), err))
		return
	}

	orgID, err := influxdb.IDFromString(reqBody.OrgID)
	if err != nil {
		s.api.Err(w, r, &influxdb.Error{
			Code: influxdb.EInvalid,
			Msg:  fmt.Sprintf("invalid organization ID provided: %q", reqBody.OrgID),
		})
		return
	}

	var stackID influxdb.ID
	if reqBody.StackID != nil {
		if err := stackID.DecodeFromString(*reqBody.StackID); err != nil {
			s.api.Err(w, r, &influxdb.Error{
				Code: influxdb.EInvalid,
				Msg:  fmt.Sprintf("invalid stack ID provided: %q", *reqBody.StackID),
			})
			return
		}
	}

	parsedPkg, err := reqBody.Pkgs(encoding)
	if err != nil {
		s.api.Err(w, r, &influxdb.Error{
			Code: influxdb.EUnprocessableEntity,
			Err:  err,
		})
		return
	}

	actions, err := reqBody.validActions()
	if err != nil {
		s.api.Err(w, r, err)
		return
	}

	applyOpts := []ApplyOptFn{
		ApplyWithEnvRefs(reqBody.EnvRefs),
		ApplyWithPkg(parsedPkg),
		ApplyWithStackID(stackID),
	}
	for _, a := range actions.SkipResources {
		applyOpts = append(applyOpts, ApplyWithResourceSkip(a))
	}
	for _, a := range actions.SkipKinds {
		applyOpts = append(applyOpts, ApplyWithKindSkip(a))
	}

	auth, err := pctx.GetAuthorizer(r.Context())
	if err != nil {
		s.api.Err(w, r, err)
		return
	}
	userID := auth.GetUserID()

	if reqBody.DryRun {
		impact, err := s.svc.DryRun(r.Context(), *orgID, userID, applyOpts...)
		if IsParseErr(err) {
			s.api.Respond(w, r, http.StatusUnprocessableEntity, RespApply{
				Sources: append([]string{}, impact.Sources...), // guarantee non nil slice
				StackID: impact.StackID.String(),
				Diff:    impact.Diff,
				Summary: impact.Summary,
				Errors:  convertParseErr(err),
			})
			return
		}
		if err != nil {
			s.api.Err(w, r, err)
			return
		}

		s.api.Respond(w, r, http.StatusOK, RespApply{
			Sources: append([]string{}, impact.Sources...), // guarantee non nil slice
			StackID: impact.StackID.String(),
			Diff:    impact.Diff,
			Summary: impact.Summary,
		})
		return
	}

	applyOpts = append(applyOpts, ApplyWithSecrets(reqBody.Secrets))

	impact, err := s.svc.Apply(r.Context(), *orgID, userID, applyOpts...)
	if err != nil && !IsParseErr(err) {
		s.api.Err(w, r, err)
		return
	}

	s.api.Respond(w, r, http.StatusCreated, RespApply{
		Sources: append([]string{}, impact.Sources...), // guarantee non nil slice
		StackID: impact.StackID.String(),
		Diff:    impact.Diff,
		Summary: impact.Summary,
		Errors:  convertParseErr(err),
	})
}

type encoder interface {
	Encode(interface{}) error
}

func formatSources(sources []string) string {
	return strings.Join(sources, "; ")
}

func decodeWithEncoding(r *http.Request, v interface{}) (Encoding, error) {
	encoding := pkgEncoding(r.Header.Get("Content-Type"))

	var dec interface{ Decode(interface{}) error }
	switch encoding {
	case EncodingJsonnet:
		dec = jsonnet.NewDecoder(r.Body)
	case EncodingYAML:
		dec = yaml.NewDecoder(r.Body)
	default:
		dec = json.NewDecoder(r.Body)
	}

	return encoding, dec.Decode(v)
}

func pkgEncoding(contentType string) Encoding {
	switch contentType {
	case "application/x-jsonnet":
		return EncodingJsonnet
	case "text/yml", "application/x-yaml":
		return EncodingYAML
	default:
		return EncodingJSON
	}
}

func convertEncoding(ct, rawURL string) Encoding {
	ct = strings.ToLower(ct)
	urlBase := path.Ext(rawURL)
	switch {
	case ct == "jsonnet" || urlBase == ".jsonnet":
		return EncodingJsonnet
	case ct == "json" || urlBase == ".json":
		return EncodingJSON
	case ct == "yml" || ct == "yaml" || urlBase == ".yml" || urlBase == ".yaml":
		return EncodingYAML
	default:
		return EncodingSource
	}
}

func newJSONEnc(w io.Writer) encoder {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	return enc
}

func (s *HTTPServer) encResp(w http.ResponseWriter, r *http.Request, enc encoder, code int, res interface{}) {
	w.WriteHeader(code)
	if err := enc.Encode(res); err != nil {
		s.api.Err(w, r, &influxdb.Error{
			Msg:  fmt.Sprintf("unable to marshal; Err: %v", err),
			Code: influxdb.EInternal,
			Err:  err,
		})
	}
}

func convertParseErr(err error) []ValidationErr {
	pErr, ok := err.(ParseError)
	if !ok {
		return nil
	}
	return pErr.ValidationErrs()
}

func newDecodeErr(encoding string, err error) *influxdb.Error {
	return &influxdb.Error{
		Msg:  fmt.Sprintf("unable to unmarshal %s", encoding),
		Code: influxdb.EInvalid,
		Err:  err,
	}
}

func convertStackToRespStack(st Stack) RespStack {
	resources := make([]RespStackResource, 0, len(st.Resources))
	for _, r := range st.Resources {
		asses := make([]RespStackResourceAssoc, 0, len(r.Associations))
		for _, a := range r.Associations {
			asses = append(asses, RespStackResourceAssoc{
				Kind:     a.Kind,
				MetaName: a.MetaName,
			})
		}
		resources = append(resources, RespStackResource{
			APIVersion:   r.APIVersion,
			ID:           r.ID.String(),
			Kind:         r.Kind,
			MetaName:     r.MetaName,
			Associations: asses,
		})
	}

	return RespStack{
		ID:          st.ID.String(),
		OrgID:       st.OrgID.String(),
		Name:        st.Name,
		Description: st.Description,
		Resources:   resources,
		Sources:     append([]string{}, st.Sources...),
		URLs:        append([]string{}, st.URLs...),
		CRUDLog:     st.CRUDLog,
	}
}
