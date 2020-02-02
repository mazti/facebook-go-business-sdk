package sdk

type APIResponse interface {
	GetBody() []byte
	GetHeader() []byte
	SetBody(body []byte)
	SetHeader(header []byte)
}
