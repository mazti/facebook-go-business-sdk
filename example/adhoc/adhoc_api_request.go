package main

import (
	"fmt"

	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
)

func Log(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	context := sdk.NewContext(
		sdk.AppSecret(AppSecret),
		sdk.AccessToken(AccessToken),
		sdk.Logger(Log),
		sdk.Debug(true),
	)
	response := sdk.APINodeList{
		Data: []sdk.AdAccount{},
	}

	req := sdk.NewAPIRequest(
		context,
		"me",
		"adaccounts",
		"GET",
		sdk.ResponseTemplate(response),
	)

	resp, err := req.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nodeList := resp.(sdk.APINodeList)

	fmt.Println(nodeList.Data.([]sdk.AdAccount))
}
