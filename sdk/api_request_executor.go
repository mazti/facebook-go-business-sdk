package sdk

import (
	"errors"
	"net/http"
)

type IRequestExecutor interface {
	Execute(method, apiUrl string, params map[string]interface{}, context APIContext) (ResponseWrapper, error)
	SendGet(apiUrl, params map[string]interface{}, context APIContext) (ResponseWrapper, error)
	SendPost(apiUrl, params map[string]interface{}, context APIContext) (ResponseWrapper, error)
	SendDelete(apiUrl, params map[string]interface{}, context APIContext) (ResponseWrapper, error)
}

type DefaultRequestExecutor struct {
}

func (e *DefaultRequestExecutor) Execute(
	method, apiUrl string, params map[string]interface{}, context APIContext) (resp ResponseWrapper, err error) {
	if http.MethodGet == method {

	}
	if http.MethodPost == method {

	}
	if http.MethodDelete == method {

	}
	return resp, errors.New("unsupported method")
}

func (e *DefaultRequestExecutor) SendGet(
	apiUrl, params map[string]interface{}, context APIContext) (resp ResponseWrapper, err error) {

	return
}

func (e *DefaultRequestExecutor) SendPost(
	apiUrl, params map[string]interface{}, context APIContext) (resp ResponseWrapper, err error) {
	return
}

func (e *DefaultRequestExecutor) SendDelete(
	apiUrl, params map[string]interface{}, context APIContext) (resp ResponseWrapper, err error) {
	return
}
