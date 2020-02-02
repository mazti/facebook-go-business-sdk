package sdk

type APINode struct {
	Context *APIContext
	Body    []byte
	Header  []byte
}

func CreateAPINode(context *APIContext) *APINode {
	return &APINode{
		Context: context,
	}
}

func LoadJSON(ctx *APIContext, body []byte, header []byte) *APINode {
	return &APINode{
		Context: ctx,
		Body:    body,
		Header:  header,
	}
}

func (n APINode) GetBody() []byte {
	return n.Body
}

func (n APINode) GetHeader() []byte {
	return n.Header
}
