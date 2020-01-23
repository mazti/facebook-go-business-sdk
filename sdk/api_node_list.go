package sdk

type Cursors struct {
	Before string `json:"before,omitempty"`
	After  string `json:"after,omitempty"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
	Next    string  `json:"next,omitempty"`
}

type APINodeList struct {
	Paging         Paging `json:"paging"`
	request        *APIRequest
	body           []byte
	header         []byte
	autoPagination bool
	appSecret      string
}
