package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

const (
	defaultTimeout = 15
	userAgent      = "User-Agent"
	contentType    = "Content-Type"
	contentTypeURL = "application/x-www-form-urlencoded"
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
	context.Log("Request:")
	context.Log("GET:", apiURL)
	httpReq, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add(userAgent, context.userAgent)
	httpReq.Header.Add(contentType, contentTypeURL)
	httpResp, err := e.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return readResponse(httpResp)
}

func (e *DefaultRequestExecutor) SendPost(apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	apiURL, err = createRequestURL(apiURL, params)
	if err != nil {
		return nil, err
	}
	context.Log("Request:")
	context.Log("POST:", apiURL)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val.(string))
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add(userAgent, context.userAgent)
	httpResp, err := e.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return readResponse(httpResp)
}

func (e *DefaultRequestExecutor) SendDelete(apiURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	apiURL, err = createRequestURL(apiURL, params)
	if err != nil {
		return nil, err
	}
	context.Log("Request:")
	context.Log("DELETE:", apiURL)
	httpReq, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Add(userAgent, context.userAgent)
	httpResp, err := e.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	return readResponse(httpResp)
}

func readResponse(resp *http.Response) (*ResponseWrapper, error) {
	defer func() { _ = resp.Body.Close() }()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respHeader, err := json.Marshal(resp.Header)
	if err != nil {
		return nil, err
	}
	return &ResponseWrapper{
		Body:   respBody,
		Header: respHeader,
	}, nil
}
