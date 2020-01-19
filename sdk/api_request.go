package sdk

type APIRequest struct {
	context          *APIContext
	executor         IRequestExecutor
	userAgent        string
	useVideoEndpoint bool
	nodeID           string
	endpoint         string
	method           string
	paramNames       []string
	params           map[string]interface{}
	returnFields     []string
	overrideURL      string
	lastResponse     APIResponse
}

type ResponseWrapper struct {
	body   []byte
	header []byte
}

func NewResponseWrapper(body []byte, header []byte) *ResponseWrapper {
	return &ResponseWrapper{
		body:   body,
		header: header,
	}
}

func (req *APIRequest) Execute() (APIResponse, error) {
	return req.ExecuteWithParams(nil)
}

func (req *APIRequest) ExecuteWithParams(extraParams map[string]interface{}) (APIResponse, error) {
	rw := req.executeInternal(extraParams)
	req.lastResponse = req.parseResponse(rw.body, rw.header)
	return req.lastResponse, nil
}

func (req *APIRequest) executeInternal(extraParams map[string]interface{}) *ResponseWrapper {

	return &ResponseWrapper{
		body:   nil,
		header: nil,
	}
}

func (req *APIRequest) parseResponse(body []byte, header []byte) APIResponse {
	return nil
}
