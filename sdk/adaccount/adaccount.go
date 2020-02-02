package adaccount

import (
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
	adAccount := &AdAccount{
		node: sdk.CreateAPINode(context),
		ID:   id,
	}
	return adAccount
}

func (ent *AdAccount) getPrefixID() string {
	return fmt.Sprintf(prefix, ent.ID)
}

func (ent *AdAccount) Fetch() (*AdAccount, error) {
	obj, err := FetchByID(ent.getPrefixID(), ent.node.Context)
	if err != nil {
		return nil, err
	}
	// TODO: check copy value
	return obj, nil
}

func FetchByID(id string, context *sdk.APIContext) (*AdAccount, error) {
	req := sdk.NewAPIRequest(
		context,
		id,
		sdk.DefaultEndpoint,
		http.MethodGet,
		sdk.Parser(sdk.ParserResponse),
	)
	_, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
