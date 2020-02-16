package adaccount

type AccountStatus int64

const (
	Active                 AccountStatus = 1
	Disable                AccountStatus = 2
	Unsettled              AccountStatus = 3
	PendingRiskReview      AccountStatus = 7
	PendingSettlement      AccountStatus = 8
	InGracePeriod          AccountStatus = 9
	PendingClosure         AccountStatus = 100
	Closed                 AccountStatus = 101
	AnyActive              AccountStatus = 201
	AnyClosedAccountStatus AccountStatus = 202
)
