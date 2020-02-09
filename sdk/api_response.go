package sdk

type APIResponse interface {
	SetContext(*APIContext)
	Load(*APIContext, *APIRequest, []byte, []byte)
}
