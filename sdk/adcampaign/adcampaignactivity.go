package adcampaign

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mazti/facebook-go-business-sdk/sdk"
)

type AdCampaignActivity struct {
	node          *sdk.APINode
	ID            string `json:"id"`
}

func NewAdCampaignActivity(id string, context *sdk.APIContext) *AdCampaignActivity {
	return &AdCampaignActivity{
		node: sdk.CreateAPINode(context),
		ID:   id,
	}
}
func ParserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdCampaignActivity{}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}

func FetchByID(id string, context *sdk.APIContext) (*AdCampaignActivity, error) {
	req := sdk.NewAPIRequest(
		context,
		id,
		sdk.DefaultEndpoint,
		http.MethodGet,
		sdk.Parser(ParserResponse),
	)
	ent, err := req.Execute()
	if err != nil {
		return nil, err
	}

	account := ent.(*AdCampaignActivity)

	return account, nil
}

func (ent *AdCampaignActivity) Fetch() (*AdCampaignActivity, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.node.Context)
	if err != nil {
		return nil, err
	}
	// TODO: check copy value
	return obj, nil
}

func (ent *AdCampaignActivity) getPrefixID() string {
	return fmt.Sprintf(prefix, ent.ID)
}