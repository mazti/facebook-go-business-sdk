package main

import (
	"encoding/json"
	"fmt"

	. "github.com/mazti/facebook-go-business-sdk/example"
	"github.com/mazti/facebook-go-business-sdk/sdk"
	"github.com/mazti/facebook-go-business-sdk/sdk/adaccount"
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
	req := sdk.NewAPIRequest(
		context,
		"me",
		"adaccounts",
		"GET",
		sdk.Parser(sdk.ParseAPINodeList),
	)
	resp, err := req.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	nodeList := resp.(sdk.APINodeList)

	fmt.Println("-----")
	fmt.Println(string(nodeList.Data))

	fmt.Println("-----")
	var adaccounts []adaccount.AdAccount

	err = json.Unmarshal(nodeList.Data, &adaccounts)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(adaccounts)
}
