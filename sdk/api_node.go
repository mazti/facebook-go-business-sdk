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

func LoadJSON(ctx *APIContext, body []byte, header []byte) *APINode {
	return &APINode{
		Context: ctx,
		Body:    body,
		Header:  header,
	}
}

func (ent APINode) GetBody() []byte {
	return ent.Body
}

func (ent APINode) GetHeader() []byte {
	return ent.Header
}

func (ent *APINode) SetRequest(request *APIRequest) {
	ent.Request = request
}

func (ent *APINode) SetBody(body []byte) {
	ent.Body = body
}

func (ent *APINode) SetHeader(header []byte) {
	ent.Header = header
}
