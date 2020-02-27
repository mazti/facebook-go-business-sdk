package businessuser

type Role string

const (
	FinanceEditor     Role = "FINANCE_EDITOR"
	FinanceAnalyst    Role = "FINANCE_ANALYST"
	AdsRightsReviewer Role = "ADS_RIGHTS_REVIEWER"
	Admin             Role = "ADMIN"
	Employee          Role = "EMPLOYEE"
)
