package sdk

import "encoding/json"

type APINode struct {
	Context  *APIContext
	RawValue []byte
	Header   []byte
}

func CreateAPINode(context *APIContext) *APINode {
	return &APINode{
		Context: context,
	}
}

func LoadJSON(ctx *APIContext, body []byte, header []byte) *APINode {
	return &APINode{
		Context:  ctx,
		RawValue: body,
		Header:   header,
	}
}

func (n APINode) GetRawResponse() []byte {
	return n.RawValue
}

func (n APINode) GetRawResponseAsJsonObject() json.RawMessage {
	return n.RawValue
}

func (n APINode) GetHeader() []byte {
	return n.Header
}
