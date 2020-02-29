package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	UnsupportedResponse = errors.New("unsupported response")
	ReachToTheEnd       = errors.New("reach to the end")
)

var (
	UnsupportedMethod = func(method string) error {
		return fmt.Errorf("unsupported method: %s", method)
	}
	ErrorInResponse = func(errResp *FbErrorResponse) error {
		if message, err := json.Marshal(errResp); err == nil {
			return errors.New(string(message))
		}
		return errors.New("there is an error in response")
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
