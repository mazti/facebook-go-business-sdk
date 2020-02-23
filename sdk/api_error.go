package sdk

import (
	"errors"
	"fmt"
)

var (
	UnsupportedResponse = errors.New("unsupported response")
	ReachToTheEnd       = errors.New("reach to the end")
	ErrorInResponse     = errors.New("there is an error in response")
)

var (
	UnsupportedMethod = func(method string) error {
		return fmt.Errorf("unsupported method: %s", method)
	}
)

type FbErrorMessage struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int    `json:"code"`
	FbTraceID string `json:"fbtrace_id"`
}

type FbErrorResponse struct {
	request *APIRequest
	Error   FbErrorMessage `json:"error"`
}

func (ent *FbErrorResponse) SetRequest(request *APIRequest) {
	ent.request = request
}

func (ent *FbErrorResponse) GetRequest() *APIRequest {
	return ent.request
}
