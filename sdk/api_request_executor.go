package sdk

import (
	"errors"
	"net/http"
	"time"
)

const (
	defaultTimeout = 15
)

type IRequestExecutor interface {
	Execute(method, url string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendGet(url string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendPost(url string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
	SendDelete(url string, params map[string]interface{}, context *APIContext) (*ResponseWrapper, error)
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

func (e *DefaultRequestExecutor) Execute(
	method, url string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	if http.MethodGet == method {
		return e.SendGet(url, params, context)
	}
	if http.MethodPost == method {
		return e.SendPost(url, params, context)
	}
	if http.MethodDelete == method {
		return e.SendDelete(url, params, context)
	}
	return resp, errors.New("unsupported method")
}

func (e *DefaultRequestExecutor) SendGet(
	apURL string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	//url.
	return
}

func (e *DefaultRequestExecutor) SendPost(
	url string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	return
}

func (e *DefaultRequestExecutor) SendDelete(
	url string, params map[string]interface{}, context *APIContext) (resp *ResponseWrapper, err error) {
	return
}
