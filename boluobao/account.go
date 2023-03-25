package boluobao

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao/boluobaoapi"
)

func ShowUserInformation() {
	response := boluobaoapi.AccountInformationAPI()
	fmt.Println("用户名:", response.Get("data.nickName").String())
	fmt.Println("用户ID:", response.Get("data.accountId").String())

}

func UserLogin(username string, password string, retry int) {
	if retry <= 0 {
		retry = 1
	} else if retry > 3 {
		retry = 3
	}
	for i := 0; i < retry; i++ {
		response := boluobaoapi.LoginAccountAPI(username, password)
		if response != "" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Println("retry login account:", i+1)
		}
	}
}
