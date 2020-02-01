package adaccount

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
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

func (ent *AdAccount) Fetch() (*AdAccount, error) {
	return FetchByID(ent.ID, ent.node.Context)
}

func FetchByID(id string, context *sdk.APIContext) (*AdAccount, error) {
	return nil, nil
}
