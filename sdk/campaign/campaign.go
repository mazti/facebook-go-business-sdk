package campaign

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

type Campaign struct {
	node        *sdk.APINode
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	AdLabels    string `json:"adlabels"`
	BidStrategy string `json:"bid_strategy"`
}
