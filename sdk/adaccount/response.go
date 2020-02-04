package adaccount

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func (ent AdAccount) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent AdAccount) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *AdAccount) SetRequest(request *sdk.APIRequest) {
	ent.node.SetRequest(request)
}

func (ent *AdAccount) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *AdAccount) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
