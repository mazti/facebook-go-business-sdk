package sdk

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	accessTokenKey    = "access_token"
	appSecretProofKey = "appsecret_proof"
	fieldsKey         = "fields"
	limitKey          = "limit"
	afterKey          = "after"
)

// ResponseWrapper ...
type ResponseWrapper struct {
	Status     string
	StatusCode int
	Body       []byte
	Header     []byte
}

// Unmarshal ...
type Unmarshal func(json json.RawMessage) (APIResponse, error)

// APIRequest ...
type APIRequest struct {
	Context *APIContext

	executor  IRequestExecutor
	unmarshal Unmarshal

	nodeID   string
	endpoint string
	method   string

	paramNames   []string
	params       map[string]interface{}
	returnFields []string
	overrideURL  string

	useVideoEndpoint bool
	userAgent        string

	lastResponse APIResponse
}

func NewAPIRequest(context *APIContext, nodeID, endpoint, method string, options ...func(*APIRequest)) *APIRequest {
	req := &APIRequest{
		Context:   context,
		nodeID:    nodeID,
		endpoint:  endpoint,
		method:    method,
		executor:  NewDefaultRequestExecutor(),
		unmarshal: ParserResponse,
	}

	for _, option := range options {
		option(req)
	}

	return req
}

// Options for constructors
func Parser(unmarshal func(json json.RawMessage) (APIResponse, error)) func(*APIRequest) {
	return func(req *APIRequest) {
		req.unmarshal = unmarshal
	}
}

func ReturnFields(fields []string) func(*APIRequest) {
	return func(req *APIRequest) {
		req.returnFields = fields
	}
}

func ParamNames(params []string) func(*APIRequest) {
	return func(req *APIRequest) {
		req.paramNames = params
	}
}

func (req *APIRequest) Execute() (APIResponse, error) {
	return req.ExecuteWithParams(nil)
}

func (req *APIRequest) ExecuteWithParams(extraParams map[string]interface{}) (APIResponse, error) {
	rw, err := req.executeInternal(extraParams)
	if err != nil {
		return nil, err
	}
	req.lastResponse = req.parseResponse(rw.Body, rw.Header)
	return req.lastResponse, nil
}

func (req *APIRequest) SetOverrideURL(url string) {
	req.overrideURL = url
}

func (req *APIRequest) executeInternal(extraParams map[string]interface{}) (*ResponseWrapper, error) {
	ctx := req.Context
	ctx.Log("========Start of API Call========")
	defer ctx.Log("========End of API Call========")
	resp, err := req.executor.Execute(req.method, req.getURL(), req.getAllParams(extraParams), req.Context)
	if err != nil {
		return nil, err
	}
	ctx.Log("Response:", string(resp.Body))
	return resp, nil
}

func (req *APIRequest) parseResponse(body []byte, header []byte) APIResponse {
	// TODO: parse error message {"error":{"message":"Malformed access token <token>","type":"OAuthException","code":190,"fbtrace_id":"AI3xqjaZLDAl4hv6Ng7KfZ8"}}
	if req.unmarshal != nil {
		resp, err := req.unmarshal(body)
		if err == nil && resp != nil {
			resp.SetRequest(req)
			return resp
		}
	}
	return Load(
		req,
		body,
		header,
	)
}

func (req *APIRequest) getURL() string {
	if len(req.overrideURL) > 0 {
		return req.overrideURL
	}
	ctx := req.Context
	endpointBase := ctx.endpointBase
	if req.useVideoEndpoint {
		endpointBase = ctx.videoEndpointBase
	}
	return fmt.Sprintf("%s/%s/%s/%s", endpointBase, ctx.version, req.nodeID, req.endpoint)
}

func (req *APIRequest) getAllParams(extraParams map[string]interface{}) (allParams map[string]interface{}) {
	allParams = map[string]interface{}{}
	ctx := req.Context
	if len(req.overrideURL) > 0 {
		return allParams
	}
	for k, v := range req.params {
		allParams[k] = v
	}
	for k, v := range extraParams {
		allParams[k] = v
	}
	allParams[accessTokenKey] = ctx.accessToken
	if len(ctx.appSecret) > 0 {
		allParams[appSecretProofKey] = ctx.GetAppSecretProof()
	}
	if len(req.returnFields) > 0 {
		allParams[fieldsKey] = strings.Join(req.returnFields, ",")
	}
	return allParams
}
