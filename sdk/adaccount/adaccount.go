package adaccount

import (
	"encoding/json"
	"fmt"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adsinsights"
	"github.com/mazti/facebook-go-business-sdk/sdk/campaign"
	"net/http"
	"strings"
)

type AdAccount struct {
	Context       *sdk.APIContext
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	AccountStatus int64  `json:"account_status"`
	// TODO: adding more field here
}

func NewAdAccount(id string, context *sdk.APIContext) *AdAccount {
	return &AdAccount{
		Context: context,
		ID:      id,
	}
}

func FetchByID(id string, context *sdk.APIContext) (*AdAccount, error) {
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

	account := ent.(*AdAccount)
	account.ID = strings.Replace(account.ID, "act_", "", 1)

	return account, nil
}

func (ent *AdAccount) Fetch() (*AdAccount, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.Context)
	if err != nil {
		return nil, err
	}
	ent.copy(obj)
	return ent, nil
}

func (ent *AdAccount) GetCampaigns() (*sdk.APINodeList, error) {
	req := campaign.CreateGetCampaignsRequest(ent.getPrefixID(), ent.Context)
	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}

func (ent *AdAccount) GetInsights() (*adsinsights.AdsInsights, error) {
	req := adsinsights.CreateGetAdsInsightsRequest(ent.getPrefixID(), ent.Context)
	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*adsinsights.AdsInsights), nil
}

//
// For internal usage
//
func (ent *AdAccount) copy(other *AdAccount) {
	ent.ID = other.ID
	ent.AccountStatus = other.AccountStatus
	ent.AccountID = other.AccountID
}
func (ent *AdAccount) getPrefixID() string {
	return fmt.Sprintf(prefix, ent.ID)
}

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdAccount{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
