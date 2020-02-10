package sdk

type APINode struct {
	Context *APIContext
	Request *APIRequest
	Body    []byte
	Header  []byte
}

func (ent *APINode) SetContext(context *APIContext) {
	ent.Context = context
}

func (ent *APINode) GetContext() *APIContext {
	return ent.Context
}

func Load(context *APIContext, req *APIRequest, body []byte, header []byte) *APINode {
	return &APINode{
		Context: context,
		Request: req,
		Body:    body,
		Header:  header,
	}
}
