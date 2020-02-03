package campaign

import (
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/campaigns"
)

type Campaign struct {
	node        *sdk.APINode
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	AdLabels    string `json:"adlabels"`
	BidStrategy string `json:"bid_strategy"`
	// TODO: adding more field here
}

func CreateGetCampaignsRequest(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.RequestParams(params),
		sdk.RequestFields(fields),
	)
}
