package adaccount

import (
	"encoding/json"
	"fmt"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"net/http"
)

type AdAccount struct {
	node          *sdk.APINode
	ID            string `json:"id"`
	AccountID     string `json:"account_id"`
	AccountStatus int64  `json:"account_status"`
}

func NewAdAccount(id string, context *sdk.APIContext) *AdAccount {
	return &AdAccount{
		node: sdk.CreateAPINode(context),
		ID:   id,
	}
}

func ParserResponse(data json.RawMessage) (sdk.APIResponse, error) {
	ent := &AdAccount{
		node: sdk.CreateAPINode(nil),
	}
	if err := json.Unmarshal(data, ent); err != nil {
		return ent, err
	}
	return ent, nil
}

func FetchByID(id string, context *sdk.APIContext) (*AdAccount, error) {
	req := sdk.NewAPIRequest(
		context,
		id,
		sdk.DefaultEndpoint,
		http.MethodGet,
		sdk.Parser(ParserResponse),
		sdk.RequestFields(fields),
	)
	//map[string]interface{}{"page_id": "abc"}
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
	// TODO: check copy value
	return obj, nil
}

func (ent *AdAccount) getPrefixID() string {
	return fmt.Sprintf(prefix, ent.ID)
}
