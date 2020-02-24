package adaccount

type AccountStatus uint32

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

type DisableReason uint32

const (
	None                  DisableReason = 0
	AdsIntegrityPolicy    DisableReason = 1
	AdsIpReview           DisableReason = 2
	RiskPayment           DisableReason = 3
	GrayAccountShutDown   DisableReason = 4
	AdsAfcReview          DisableReason = 5
	BusinessIntegrityRar  DisableReason = 6
	PermanentClose        DisableReason = 7
	UnusedResellerAccount DisableReason = 8
	UnusedAccount         DisableReason = 9
)
