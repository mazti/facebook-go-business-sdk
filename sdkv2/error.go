package sdkv2

import (
	"net/http"
)

type FacebookRequestError struct {
	message string
	status  string
	url     string
	method  string
}

// TODO: process response
func NewFacebookRequestError(response *http.Response, method, url string) *FacebookRequestError {
	return &FacebookRequestError{
		message: "",
		status:  response.Status,
		url:     url,
		method:  method,
	}
}

func (e FacebookRequestError) Error() string {
	return e.message
}
