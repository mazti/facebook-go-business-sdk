package main

import (
	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/businessuser"
)

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)

	nodeList, err := businessuser.GetBusinessUsersOfMe(context)

	if err != nil {
		context.Log(err)
		return
	}
	for nodeList != nil {
		bisUsers, err := businessuser.ParseResponse(nodeList)
		context.Log(err)

		context.Log("biz users:", len(bisUsers))

		nodeList, err = nodeList.Next(0)
		context.Log("next err:", err)
	}
}
