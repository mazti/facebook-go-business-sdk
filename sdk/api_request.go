package sdk

import (
	"fmt"
	"strings"
)

const (
	accessTokenKey = "access_token"
	appSecretProof = "appsecret_proof"
	fields         = "fields"
)

// ResponseWrapper ...
type ResponseWrapper struct {
	Body   []byte
	Header []byte
}

// APIRequest ...
type APIRequest struct {
	context          *APIContext
	executor         IRequestExecutor
	response         APIResponse
	nodeID           string
	endpoint         string
	method           string
	paramNames       []string
	params           map[string]interface{}
	returnFields     []string
	overrideURL      string
	useVideoEndpoint bool
	userAgent        string
	lastResponse     APIResponse
}

func NewAPIRequest(context *APIContext, nodeID, endpoint, method string, options ...func(*APIRequest)) *APIRequest {
	req := &APIRequest{
		context:  context,
		nodeID:   nodeID,
		endpoint: endpoint,
		method:   method,
		executor: NewDefaultRequestExecutor(),
	}

	for _, option := range options {
		option(req)
	}

	return req
}

// Options for constructors
func ResponseTemplate(resp APIResponse) func(*APIRequest) {
	return func(req *APIRequest) {
		req.response = resp
	}
}

func (req *APIRequest) Execute() (APIResponse, error) {
	return req.ExecuteWithParams(nil)
}

func (req *APIRequest) ExecuteWithParams(extraParams map[string]interface{}) (APIResponse, error) {
	rw := req.executeInternal(extraParams)
	req.lastResponse = req.parseResponse(rw.Body, rw.Header)
	return req.lastResponse, nil
}

func (req *APIRequest) executeInternal(extraParams map[string]interface{}) *ResponseWrapper {
	ctx := req.context
	ctx.Log("========Start of API Call========")
	defer ctx.Log("========End of API Call========")
	resp, err := req.executor.Execute(req.method, req.getURL(), req.getAllParams(extraParams), req.context)
	if err != nil {
		return resp
	}
	ctx.Log("Response:", string(resp.Body))
	return resp
}

func (req *APIRequest) parseResponse(body []byte, header []byte) APIResponse {
	var err error
	if req.response != nil {
		if req.response, err = req.response.Parse(body); err == nil {
			return req.response
		}
	}
	return LoadJSON(
		req.context,
		body,
		header,
	)
}

func (req *APIRequest) getURL() string {
	if len(req.overrideURL) > 0 {
		return req.overrideURL
	}
	ctx := req.context
	endpointBase := ctx.endpointBase
	if req.useVideoEndpoint {
		endpointBase = ctx.videoEndpointBase
	}
	return fmt.Sprintf("%s/%s/%s/%s", endpointBase, ctx.version, req.nodeID, req.endpoint)
}

func (req *APIRequest) getAllParams(extraParams map[string]interface{}) (allParams map[string]interface{}) {
	allParams = map[string]interface{}{}
	ctx := req.context
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
		allParams[appSecretProof] = ctx.GetAppSecretProof()
	}
	if len(req.returnFields) > 0 {
		allParams[fields] = strings.Join(req.returnFields, ",")
	}
	return allParams
}
