package campaign

func (ent Campaign) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent Campaign) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *Campaign) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *Campaign) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
