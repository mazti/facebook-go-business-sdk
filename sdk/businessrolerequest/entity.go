package businessrolerequest

import (
	"github.com/mazti/facebook-go-business-sdk/sdk/reference"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

const (
	endpoint = "/pending_users"
)

type Entity struct {
	request *sdk.APIRequest

	CreatedBy      reference.Entity `json:"created_by"`
	CreatedTime    string           `json:"created_time"`
	Email          string           `json:"email"`
	ExpirationTime string           `json:"expiration_time"`
	ExpiryTime     string           `json:"expiry_time"`
	FinanceRole    string           `json:"finance_role"`
	Id             string           `json:"id"`
	InviteLink     string           `json:"invite_link"`
	IpRole         string           `json:"ip_role"`
	Owner          reference.Entity `json:"owner"`
	Role           string           `json:"role"`
	Status         string           `json:"status"`
	UpdatedBy      reference.Entity `json:"updated_by"`
	UpdatedTime    string           `json:"updated_time"`
}

func GetPendingUsers(nodeID string, context *sdk.APIContext) (*sdk.APINodeList, error) {
	req := sdk.NewAPIRequest(
		context,
		nodeID,
		endpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
		sdk.ParamNames(params),
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
