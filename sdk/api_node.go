package sdk

type APINode struct {
	Context *APIContext
	Request *APIRequest
	Body    []byte
	Header  []byte
}

func CreateAPINode(context *APIContext) *APINode {
	return &APINode{
		Context: context,
	}
}

func (ent *APINode) Load(context *APIContext, req *APIRequest, header []byte, body []byte) {
	ent.Context = context
	ent.Request = req
	ent.Header = header
	ent.Body = body
}

func (ent *APINode) SetContext(context *APIContext) {
	ent.Context = context
}

func Load(context *APIContext, req *APIRequest, body []byte, header []byte) *APINode {
	return &APINode{
		Context: context,
		Request: req,
		Body:    body,
		Header:  header,
	}
}
