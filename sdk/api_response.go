package sdk

type APIResponse interface {
	SetContext(*APIContext)
	GetContext() *APIContext
}
