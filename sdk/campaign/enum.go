package campaign

type Status string

const (
	Active   Status = "ACTIVE"
	Archived        = "ARCHIVED"
	Deleted         = "DELETED"
	Paused          = "PAUSED"
)

type BidStrategy string

const (
	LowestCostWithoutCap BidStrategy = "LOWEST_COST_WITHOUT_CAP"
	LowestCostWithBidCap             = "LOWEST_COST_WITH_BID_CAP"
	TargetCost                       = "TARGET_COST"
)

type ConfiguredStatus string

const (
	ConfiguredActive   ConfiguredStatus = "ACTIVE"
	ConfiguredArchived                  = "ARCHIVED"
	ConfiguredDeleted                   = "DELETED"
	ConfiguredPaused                    = "PAUSED"
)

type EffectiveStatus string

const (
	EffectiveActive     ConfiguredStatus = "ACTIVE"
	EffectiveArchived                    = "ARCHIVED"
	EffectiveDeleted                     = "DELETED"
	EffectiveInProcess                   = "IN_PROCESS"
	EffectiveWithIssues                  = "WITH_ISSUES"
)
