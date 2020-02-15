package adlabel

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *Entity) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *Entity) GetContext() *sdk.APIContext {
	return ent.Context
}
