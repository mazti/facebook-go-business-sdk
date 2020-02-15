package adset

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdSet) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdSet) GetContext() *sdk.APIContext {
	return ent.Context
}
