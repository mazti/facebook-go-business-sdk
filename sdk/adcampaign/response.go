package adcampaign

func (ent AdCampaignActivity) GetBody() []byte {
	return ent.node.GetBody()
}

func (ent AdCampaignActivity) GetHeader() []byte {
	return ent.node.GetBody()
}

func (ent *AdCampaignActivity) SetBody(body []byte) {
	ent.node.SetBody(body)
}

func (ent *AdCampaignActivity) SetHeader(header []byte) {
	ent.node.SetHeader(header)
}
