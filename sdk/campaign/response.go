package campaign

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent *Campaign) Load(context *sdk.APIContext, req *sdk.APIRequest, header []byte, body []byte) {
	ent.node.Load(context, req, header, body)
}

func (ent *Campaign) SetContext(context *sdk.APIContext) {
	ent.node.SetContext(context)
}
