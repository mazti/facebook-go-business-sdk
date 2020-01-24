package main

import (
	"encoding/json"
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
	req := sdk.NewAPIRequest(context, "me", "adaccounts", "GET")
	resp, err := req.Execute()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Testing capability of read paging
	nodeList := sdk.APINodeList{}
	if err = json.Unmarshal(resp.GetRawResponse(), &nodeList); err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Println(string(resp.GetRawResponse()))
	//fmt.Println("-----")
	//fmt.Println(nodeList)
	fmt.Println("-----")
	fmt.Println(string(nodeList.Data))

}
