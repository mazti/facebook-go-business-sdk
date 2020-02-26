package businessuser

import (
	"encoding/json"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "business_users"
)

// https://developers.facebook.com/docs/marketing-api/business-manager/get-started#users
type Entity struct {
	request *sdk.APIRequest

	Business          json.RawMessage `json:"business"`
	Email             string          `json:"email"`
	FinancePermission string          `json:"finance_permission"`
	FirstName         string          `json:"first_name"`
	ID                string          `json:"id"`
	IpPermission      string          `json:"ip_permission"`
	LastName          string          `json:"last_name"`
	MarkedForRemoval  bool            `json:"marked_for_removal"`
	Name              string          `json:"name"`
	PendingEmail      string          `json:"pending_email"`
	Role              string          `json:"role"`
	Title             string          `json:"title"`
	TwoFacStatus      string          `json:"two_fac_status"`
}

func GetBusinessUsers(nodeID string, context *sdk.APIContext) (*sdk.APINodeList, error) {
	req := sdk.NewAPIRequest(
		context,
		nodeID,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.ReturnFields(fields),
	)

	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}

func ParseResponse(rawResp sdk.APIResponse) (resp []Entity, err error) {
	request := rawResp.GetRequest()
	context := rawResp.GetRequest().Context
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
		resp[i].SetRequest(request)
	}
	return
}
