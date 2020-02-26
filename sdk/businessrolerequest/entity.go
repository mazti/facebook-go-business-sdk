package businessrolerequest

import (
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/business"
)

const (
	endpoint = "/"
)

type Entity struct {
	request *sdk.APIRequest

	CreatedBy      interface{}     `json:"created_by"`
	CreatedTime    string          `json:"created_time"`
	Email          string          `json:"email"`
	ExpirationTime string          `json:"expiration_time"`
	ExpiryTime     string          `json:"expiry_time"`
	FinanceRole    string          `json:"finance_role"`
	Id             string          `json:"id"`
	InviteLink     string          `json:"invite_link"`
	IpRole         string          `json:"ip_role"`
	Owner          business.Entity `json:"owner"`
	Role           string          `json:"role"`
	Status         string          `json:"status"`
	UpdatedBy      string          `json:"updated_by"`
	UpdatedTime    string          `json:"updated_time"`
}
