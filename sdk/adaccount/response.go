package adaccount

func (ent AdAccount) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent AdAccount) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *AdAccount) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *AdAccount) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
