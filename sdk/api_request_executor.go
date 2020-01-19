package sdk

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	defaultTimeout = 15
)

type IRequestExecutor interface {
	Execute(method, apiURL string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendGet(apiURL string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendPost(apiURL string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendDelete(apiURL string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
}

type DefaultRequestExecutor struct {
	httpClient *http.Client
}

// NewDefaultRequestExecutor will allow setting timeout later
func NewDefaultRequestExecutor() *DefaultRequestExecutor {
	return &DefaultRequestExecutor{
		httpClient: &http.Client{
			Timeout: time.Duration(defaultTimeout) * time.Second,
		},
	}
}

func (e *DefaultRequestExecutor) Execute(method, apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	if http.MethodGet == method {
		return e.SendGet(apiURL, params, context)
	}
	if http.MethodPost == method {
		return e.SendPost(apiURL, params, context)
	}
	if http.MethodDelete == method {
		return e.SendDelete(apiURL, params, context)
	}
	return resp, errors.New("unsupported method")
}

func (e *DefaultRequestExecutor) SendGet(apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	apiURL, err = createRequestURL(apiURL, params)
	if err != nil {
		return nil, err
	}
	httpResp, err := e.httpClient.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = httpResp.Body.Close() }()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	respHeader, err := json.Marshal(httpResp.Header)
	if err != nil {
		return nil, err
	}
	resp = &ResponseWrapper{
		Body:   respBody,
		Header: respHeader,
	}
	return
}

func (e *DefaultRequestExecutor) SendPost(apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	return
}

func (e *DefaultRequestExecutor) SendDelete(apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	return
}
