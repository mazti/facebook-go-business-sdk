package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/business"
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

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
