package sdk

type APIResponse interface {
	SetRequest(*APIRequest)
	GetRequest() *APIRequest
}
