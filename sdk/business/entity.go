package business

import (
	"encoding/json"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessrolerequest"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
	"github.com/mazti/facebook-go-business-sdk/sdk/page"
	"github.com/mazti/facebook-go-business-sdk/sdk/reference"
	"net/http"
)

const (
	endpoint = "businesses"
)

const (
	emailKey = "email"
	roleKey  = "role"
)

type Entity struct {
	request *sdk.APIRequest

	BlockOfflineAnalytics           bool             `json:"block_offline_analytics"`
	CreatedBy                       reference.Entity `json:"created_by"`
	CreatedTime                     string           `json:"created_time"`
	ExtendedUpdatedTime             string           `json:"extended_updated_time"`
	ID                              string           `json:"id"`
	IsHidden                        bool             `json:"is_hidden"`
	IsInstagramEnabledInFbAnalytics bool             `json:"is_instagram_enabled_in_fb_analytics"`
	Link                            string           `json:"link"`
	Name                            string           `json:"name"`
	PaymentAccountID                string           `json:"payment_account_id"`
	PrimaryPage                     page.Entity      `json:"primary_page"`
	ProfilePictureUri               string           `json:"profile_picture_uri"`
	TimezoneID                      int64            `json:"timezone_id"`
	TwoFactorType                   string           `json:"two_factor_type"`
	UpdatedBy                       reference.Entity `json:"updated_by"`
	UpdatedTime                     string           `json:"updated_time"`
	VerificationStatus              string           `json:"verification_status"`
	Vertical                        string           `json:"vertical"`
	VerticalID                      int64            `json:"vertical_id"`
}

func FetchByID(id string, context *sdk.APIContext) (*Entity, error) {
	req := sdk.NewAPIRequest(
		context,
		id,
		sdk.DefaultEndpoint,
		http.MethodGet,
		sdk.Parser(parserResponse),
		sdk.ReturnFields(fields),
	)
	ent, err := req.Execute()
	if err != nil {
		return nil, err
	}

	business := ent.(*Entity)

	return business, nil
}

func (ent *Entity) CreateUser(email string, role businessuser.Role) (*businessuser.Entity, error) {
	params := map[string]interface{}{
		emailKey: email,
		roleKey:  string(role),
	}

	req := businessuser.NewAPIRequestCreateUser(ent.ID, ent.GetRequest().Context)
	resp, err := req.ExecuteWithParams(params)
	if err != nil {
		return nil, err
	}

	return resp.(*businessuser.Entity), nil
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
		sdk.MeNodeID,
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

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &Entity{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
