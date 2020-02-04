package sdk

type APIResponse interface {
	GetBody() []byte
	GetHeader() []byte
	SetRequest(*APIRequest)
	SetBody([]byte)
	SetHeader([]byte)
}
