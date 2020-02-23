package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/business"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
)

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	nodeList, err := business.GetBusinesses(context)

	if err != nil {
		context.Log(err)
		return
	}
	for nodeList != nil {
		businesses, err := business.ParseResponse(nodeList)
		context.Log(err)
		context.Log("businesses:", businesses)

		for _, biz := range businesses {
			nodeListUsers, err := businessuser.GetBusinessUsers(biz.ID, context)
			if err != nil {
				context.Log(err)
				continue
			}

			bizUsers, err := businessuser.ParseResponse(nodeListUsers)
			context.Log(err)

			context.Log("business users:", bizUsers)

		}

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
