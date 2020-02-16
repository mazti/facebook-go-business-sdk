package sdk

type APINode struct {
	request *APIRequest

	Body   []byte
	Header []byte
}

func (ent *APINode) SetRequest(request *APIRequest) {
	ent.request = request
}

func (ent *APINode) GetRequest() *APIRequest {
	return ent.request
}

func Load(req *APIRequest, body []byte, header []byte) *APINode {
	return &APINode{
		request: req,
		Body:    body,
		Header:  header,
	}
}
