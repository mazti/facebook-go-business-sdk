package sdk

import "encoding/json"

type APINode struct {
	context  *APIContext
	rawValue []byte
	header   []byte
}

func LoadJSON(ctx *APIContext, body []byte, header []byte) APINode {
	return APINode{
		context:  ctx,
		rawValue: body,
		header:   header,
	}
}

func (n APINode) GetRawResponse() []byte {
	return n.rawValue
}

func (n APINode) GetRawResponseAsJsonObject() json.RawMessage {
	return json.RawMessage{}
}

func (n APINode) GetHeader() []byte {
	return n.header
}
