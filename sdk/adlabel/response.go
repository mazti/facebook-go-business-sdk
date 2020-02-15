package adlabel

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdLabel) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdLabel) GetContext() *sdk.APIContext {
	return ent.Context
}
