package adstudy

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdStudy) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdStudy) GetContext() *sdk.APIContext {
	return ent.Context
}
