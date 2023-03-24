package boluobaoapi

import (
	"fmt"
	"github.com/tidwall/gjson"
)

// VerifyAPI 验证API的code是否为200
func VerifyAPI(response gjson.Result) *gjson.Result {
	if response.Get("status.httpCode").Int() == 200 {
		return &response
	} else {
		fmt.Println(response.Get("status.msg").String())
	}
	return nil

}
