package adaccount

import (
	"encoding/json"
	"fmt"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/campaign"
	"net/http"
	"strings"
)

type AdAccount struct {
	node          *sdk.APINode
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	AccountStatus int64  `json:"account_status"`
	// TODO: adding more field here
}

func NewAdAccount(id string, context *sdk.APIContext) *AdAccount {
	return &AdAccount{
		node: sdk.CreateAPINode(context),
		ID:   id,
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

	return account, nil
}

func (ent *AdAccount) Fetch() (*AdAccount, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.node.Context)
	if err != nil {
		return nil, err
	}
	obj.ID = strings.Replace(obj.ID, "act_", "", 1)
	// TODO: check copy value
	obj.node = ent.node
	return obj, nil
}

func (ent *AdAccount) GetCampaigns() (*sdk.APINodeList, error) {
	req := campaign.CreateGetCampaignsRequest(ent.getPrefixID(), ent.node.Context)
	resp, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.(*sdk.APINodeList), nil
}

//
// For internal usage
//
func (ent *AdAccount) getPrefixID() string {
	return fmt.Sprintf(prefix, ent.ID)
}

func parserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdAccount{
		node: sdk.CreateAPINode(nil),
	}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}
