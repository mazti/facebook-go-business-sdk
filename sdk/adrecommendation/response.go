package adrecommendation

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdRecommendation) SetContext(context *sdk.APIContext) {
	ent.Context = context
}

func (ent *AdRecommendation) GetContext() *sdk.APIContext {
	return ent.Context
}
