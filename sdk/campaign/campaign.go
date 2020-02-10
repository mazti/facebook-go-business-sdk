package campaign

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"net/http"
)

const (
	endpoint = "/campaigns"
)

type Campaign struct {
	Context     *sdk.APIContext
	ID          string `json:"id"`
	AccountID   string `json:"account_id"`
	AdLabels    string `json:"adlabels"`
	BidStrategy string `json:"bid_strategy"`
	// TODO: adding more field here
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Campaign, err error) {
	context := rawResp.GetContext()
	nodeList, ok := rawResp.(*sdk.APINodeList)
	if !ok {
		return nil, sdk.UnsupportedResponse
	}
	err = nodeList.Unmarshal(&resp)
	if err != nil {
		context.Log(err)
		return
	}
	for i := 0; i < len(resp); i++ {
		resp[i].SetContext(context)
	}
	return
}

func CreateGetCampaignsRequest(id string, context *sdk.APIContext) *sdk.APIRequest {
	return sdk.NewAPIRequest(
		context,
		id,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.ParamNames(params),
		sdk.ReturnFields(fields),
	)
}
