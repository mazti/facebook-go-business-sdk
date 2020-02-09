package adaccount

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *AdAccount) Load(context *sdk.APIContext, req *sdk.APIRequest, header []byte, body []byte) {
	ent.Context = context
}

func (ent *AdAccount) SetContext(context *sdk.APIContext) {
	ent.Context = context
}
