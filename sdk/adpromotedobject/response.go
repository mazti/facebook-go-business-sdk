package adpromotedobject

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdPromotedObject) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdPromotedObject) GetContext() *sdk.APIContext {
	return ent.Context
}
