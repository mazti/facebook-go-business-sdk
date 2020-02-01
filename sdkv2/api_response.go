package sdkv2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	errorKey           = "error"
	successKey         = "success"
	serviceUnavailable = "Service Unavailable"
)

type APIResponse struct {
	bodyRaw        []byte
	body           map[string]interface{}
	httpStatus     string
	httpStatusCode int
	headers        []byte
	request        *http.Request
	response       *http.Response
}

func NewAPIResponse(response *http.Response, request *http.Request) (*APIResponse, error) {
	defer func() { _ = response.Body.Close() }()
	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respHeader, err := json.Marshal(response.Header)
	if err != nil {
		return nil, err
	}
	body := map[string]interface{}{}

	err = json.Unmarshal(respBody, &body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{
		request:        request,
		response:       response,
		httpStatus:     response.Status,
		httpStatusCode: response.StatusCode,
		bodyRaw:        respBody,
		body:           body,
		headers:        respHeader,
	}, nil
}

func (resp *APIResponse) IsSuccess() bool {
	if _, ok := resp.body[errorKey]; ok {
		return false
	}
	if len(resp.body) > 0 {
		if val, ok := resp.body[successKey]; ok {
			if success, cOK := val.(bool); cOK {
				return success
			}
		}
		_, ok := resp.body[serviceUnavailable]
		return !ok
	}
	if resp.httpStatusCode == http.StatusNotModified {
		// ETag Hit
		return true
	}
	if resp.httpStatusCode == http.StatusOK {
		// HTTP OK
		return true
	}
	return false
}

func (resp *APIResponse) Error() error {
	if resp.IsSuccess() {
		return nil
	}

	return NewFacebookRequestError(
		resp.response,
		resp.request.Method,
		resp.request.URL.String(),
	)
}
