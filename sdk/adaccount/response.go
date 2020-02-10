package adaccount

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdAccount) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdAccount) GetContext() *sdk.APIContext {
	return ent.Context
}
