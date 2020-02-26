package business

import (
	"encoding/json"
	"github.com/mazti/facebook-go-business-sdk/sdk/businesscreateuser"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessrolerequest"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
	"github.com/mazti/facebook-go-business-sdk/sdk/page"
)

const (
	nodeID   = "me"
	endpoint = "businesses"
)

type Entity struct {
	request *sdk.APIRequest

	BlockOfflineAnalytics           bool            `json:"block_offline_analytics"`
	CreatedBy                       json.RawMessage `json:"created_by"`
	CreatedTime                     string          `json:"created_time"`
	ExtendedUpdatedTime             string          `json:"extended_updated_time"`
	ID                              string          `json:"id"`
	IsHidden                        bool            `json:"is_hidden"`
	IsInstagramEnabledInFbAnalytics bool            `json:"is_instagram_enabled_in_fb_analytics"`
	Link                            string          `json:"link"`
	Name                            string          `json:"name"`
	PaymentAccountID                string          `json:"payment_account_id"`
	PrimaryPage                     page.Entity     `json:"primary_page"`
	ProfilePictureUri               string          `json:"profile_picture_uri"`
	TimezoneID                      int64           `json:"timezone_id"`
	TwoFactorType                   string          `json:"two_factor_type"`
	UpdatedBy                       json.RawMessage `json:"updated_by"`
	UpdatedTime                     string          `json:"updated_time"`
	VerificationStatus              string          `json:"verification_status"`
	Vertical                        string          `json:"vertical"`
	VerticalID                      int64           `json:"vertical_id"`
}

func (ent *Entity) CreateUser(email string, role businessuser.Role) (*businesscreateuser.Entity, error) {
	resp, err := businesscreateuser.CreateAPIRequestCreateBusinessUser(ent.ID, ent.GetRequest().Context).
		SetEmail(email).
		SetRole(role).
		Execute()

	if err != nil {
		return nil, err
	}
	return resp, nil
}


func (ent *Entity) GetUsers() (*sdk.APINodeList, error) {
	resp, err := businessuser.GetBusinessUsers(ent.ID, ent.GetRequest().Context)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ent *Entity) GetPendingUsers() (*sdk.APINodeList, error) {
	resp, err := businessrolerequest.GetPendingUsers(ent.ID, ent.GetRequest().Context)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetBusinesses(context *sdk.APIContext) (*sdk.APINodeList, error) {
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
