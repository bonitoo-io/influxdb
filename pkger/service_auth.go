package pkger

import (
	"context"

	"github.com/influxdata/influxdb/v2"
)

type AuthAgent interface {
	IsWritable(ctx context.Context, orgID influxdb.ID, resType influxdb.ResourceType) error
	OrgPermissions(ctx context.Context, orgID influxdb.ID, action influxdb.Action, rest ...influxdb.Action) error
}

type authMW struct {
	authAgent AuthAgent
	next      SVC
}

var _ SVC = (*authMW)(nil)

// MWAuth is an auth service middleware for the packager domain.
func MWAuth(authAgent AuthAgent) SVCMiddleware {
	return func(svc SVC) SVC {
		return &authMW{
			authAgent: authAgent,
			next:      svc,
		}
	}
}

func (s *authMW) InitStack(ctx context.Context, userID influxdb.ID, newStack Stack) (Stack, error) {
	err := s.authAgent.IsWritable(ctx, newStack.OrgID, ResourceTypeStack)
	if err != nil {
		return Stack{}, err
	}
	return s.next.InitStack(ctx, userID, newStack)
}

func (s *authMW) DeleteStack(ctx context.Context, identifiers struct{ OrgID, UserID, StackID influxdb.ID }) error {
	err := s.authAgent.IsWritable(ctx, identifiers.OrgID, ResourceTypeStack)
	if err != nil {
		return err
	}
	return s.next.DeleteStack(ctx, identifiers)
}

func (s *authMW) ExportStack(ctx context.Context, orgID, stackID influxdb.ID) (*Pkg, error) {
	err := s.authAgent.OrgPermissions(ctx, orgID, influxdb.ReadAction)
	if err != nil {
		return nil, err
	}
	return s.next.ExportStack(ctx, orgID, stackID)
}

func (s *authMW) ListStacks(ctx context.Context, orgID influxdb.ID, f ListFilter) ([]Stack, error) {
	err := s.authAgent.OrgPermissions(ctx, orgID, influxdb.ReadAction)
	if err != nil {
		return nil, err
	}
	return s.next.ListStacks(ctx, orgID, f)
}

func (s *authMW) ReadStack(ctx context.Context, id influxdb.ID) (Stack, error) {
	st, err := s.next.ReadStack(ctx, id)
	if err != nil {
		return Stack{}, err
	}

	err = s.authAgent.OrgPermissions(ctx, st.OrgID, influxdb.ReadAction)
	if err != nil {
		return Stack{}, err
	}
	return st, nil
}

func (s *authMW) UpdateStack(ctx context.Context, upd StackUpdate) (Stack, error) {
	stack, err := s.next.ReadStack(ctx, upd.ID)
	if err != nil {
		return Stack{}, err
	}

	err = s.authAgent.IsWritable(ctx, stack.OrgID, ResourceTypeStack)
	if err != nil {
		return Stack{}, err
	}
	return s.next.UpdateStack(ctx, upd)
}

func (s *authMW) CreatePkg(ctx context.Context, setters ...CreatePkgSetFn) (*Pkg, error) {
	return s.next.CreatePkg(ctx, setters...)
}

func (s *authMW) DryRun(ctx context.Context, orgID, userID influxdb.ID, opts ...ApplyOptFn) (PkgImpactSummary, error) {
	return s.next.DryRun(ctx, orgID, userID, opts...)
}

func (s *authMW) Apply(ctx context.Context, orgID, userID influxdb.ID, opts ...ApplyOptFn) (PkgImpactSummary, error) {
	return s.next.Apply(ctx, orgID, userID, opts...)
}
