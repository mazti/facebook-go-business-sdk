package sdk

type APIResponse interface {
	GetBody() []byte
	GetHeader() []byte
}
