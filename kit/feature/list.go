// Code generated by the feature package; DO NOT EDIT.

package feature

var appMetrics = MakeBoolFlag(
	"App Metrics",
	"appMetrics",
	"Bucky, Monitoring Team",
	false,
	Permanent,
	true,
)

// AppMetrics - Send UI Telementry to Tools cluster - should always be false in OSS
func AppMetrics() BoolFlag {
	return appMetrics
}

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var communityTemplates = MakeBoolFlag(
	"Community Templates",
	"communityTemplates",
	"Bucky, Johnny Steenbergen (Berg)",
	false,
	Permanent,
	true,
)

// CommunityTemplates - Replace current template uploading functionality with community driven templates
func CommunityTemplates() BoolFlag {
	return communityTemplates
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var pushDownWindowAggregateCount = MakeBoolFlag(
	"Push Down Window Aggregate Count",
	"pushDownWindowAggregateCount",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateCount - Enable Count variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateCount() BoolFlag {
	return pushDownWindowAggregateCount
}

var pushDownWindowAggregateSum = MakeBoolFlag(
	"Push Down Window Aggregate Sum",
	"pushDownWindowAggregateSum",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateSum - Enable Sum variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateSum() BoolFlag {
	return pushDownWindowAggregateSum
}

var pushDownWindowAggregateRest = MakeBoolFlag(
	"Push Down Window Aggregate Rest",
	"pushDownWindowAggregateRest",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateRest - Enable non-Count, non-Sum variants of PushDownWindowAggregateRule and PushDownWindowAggregateRule (stage 2)
func PushDownWindowAggregateRest() BoolFlag {
	return pushDownWindowAggregateRest
}

var newAuth = MakeBoolFlag(
	"New Auth Package",
	"newAuth",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewAuthPackage - Enables the refactored authorization api
func NewAuthPackage() BoolFlag {
	return newAuth
}

var sessionService = MakeBoolFlag(
	"Session Service",
	"sessionService",
	"Lyon Hill",
	false,
	Temporary,
	true,
)

// SessionService - A temporary switching system for the new session system
func SessionService() BoolFlag {
	return sessionService
}

var pushDownGroupAggregateCount = MakeBoolFlag(
	"Push Down Group Aggregate Count",
	"pushDownGroupAggregateCount",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateCount - Enable the count variant of PushDownGroupAggregate planner rule
func PushDownGroupAggregateCount() BoolFlag {
	return pushDownGroupAggregateCount
}

var pushDownGroupAggregateSum = MakeBoolFlag(
	"Push Down Group Aggregate Sum",
	"pushDownGroupAggregateSum",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateSum - Enable the sum variant of PushDownGroupAggregate planner rule
func PushDownGroupAggregateSum() BoolFlag {
	return pushDownGroupAggregateSum
}

var pushDownGroupAggregateFirst = MakeBoolFlag(
	"Push Down Group Aggregate First",
	"pushDownGroupAggregateFirst",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateFirst - Enable the first variant of PushDownGroupAggregate planner rule
func PushDownGroupAggregateFirst() BoolFlag {
	return pushDownGroupAggregateFirst
}

var pushDownGroupAggregateLast = MakeBoolFlag(
	"Push Down Group Aggregate Last",
	"pushDownGroupAggregateLast",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateLast - Enable the last variant of PushDownGroupAggregate planner rule
func PushDownGroupAggregateLast() BoolFlag {
	return pushDownGroupAggregateLast
}

var newLabels = MakeBoolFlag(
	"New Label Package",
	"newLabels",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewLabelPackage - Enables the refactored labels api
func NewLabelPackage() BoolFlag {
	return newLabels
}

var hydratevars = MakeBoolFlag(
	"New Hydrate Vars Functionality",
	"hydratevars",
	"Ariel Salem / Monitoring Team",
	false,
	Temporary,
	true,
)

// NewHydrateVarsFunctionality - Enables a minimalistic variable hydration
func NewHydrateVarsFunctionality() BoolFlag {
	return hydratevars
}

var memoryOptimizedFill = MakeBoolFlag(
	"Memory Optimized Fill",
	"memoryOptimizedFill",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedFill - Enable the memory optimized fill()
func MemoryOptimizedFill() BoolFlag {
	return memoryOptimizedFill
}

var all = []Flag{
	appMetrics,
	backendExample,
	communityTemplates,
	frontendExample,
	pushDownWindowAggregateCount,
	pushDownWindowAggregateSum,
	pushDownWindowAggregateRest,
	newAuth,
	sessionService,
	pushDownGroupAggregateCount,
	pushDownGroupAggregateSum,
	pushDownGroupAggregateFirst,
	pushDownGroupAggregateLast,
	newLabels,
	hydratevars,
	memoryOptimizedFill,
}

var byKey = map[string]Flag{
	"appMetrics":                   appMetrics,
	"backendExample":               backendExample,
	"communityTemplates":           communityTemplates,
	"frontendExample":              frontendExample,
	"pushDownWindowAggregateCount": pushDownWindowAggregateCount,
	"pushDownWindowAggregateSum":   pushDownWindowAggregateSum,
	"pushDownWindowAggregateRest":  pushDownWindowAggregateRest,
	"newAuth":                      newAuth,
	"sessionService":               sessionService,
	"pushDownGroupAggregateCount":  pushDownGroupAggregateCount,
	"pushDownGroupAggregateSum":    pushDownGroupAggregateSum,
	"pushDownGroupAggregateFirst":  pushDownGroupAggregateFirst,
	"pushDownGroupAggregateLast":   pushDownGroupAggregateLast,
	"newLabels":                    newLabels,
	"hydratevars":                  hydratevars,
	"memoryOptimizedFill":          memoryOptimizedFill,
}
