package example

import (
	"fmt"
)

var (
	AccessToken = "<access_token>"
	AppSecret   = "<app_secret>"
	AppID       = "<app_id>"
	AccountID   = "<account_id"
)

func Log(a ...interface{}) {
	fmt.Println(a...)
}
