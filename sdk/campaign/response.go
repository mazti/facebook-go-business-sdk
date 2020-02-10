package campaign

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *Campaign) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *Campaign) GetContext() *sdk.APIContext {
	return ent.Context
}
