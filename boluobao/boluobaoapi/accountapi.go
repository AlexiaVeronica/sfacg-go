package boluobaoapi

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
)

func AccountInformationAPI() *gjson.Result {
	params := map[string]string{"expand": "vipInfo,welfareCoin,isRealNameAuth,changeNickNameInfo,welfareMoney,redpacketCode,useWelfaresys,usedRedpacketCode,hasOrderChapWithFireMoney,hasUnlockChapWithAd,hasActiveUnlockChapWithAd,hasOrderedVipChaps,hasPaidFirstTime,growup,newVip"}

	return VerifyAPI(request.Get("user").AddAll(params).Json())
}

func AccountMoneyAPI() *gjson.Result {
	return VerifyAPI(request.Get("user/money").Json())
}

func AccountIPAPI() *gjson.Result {
	return VerifyAPI(request.Get("position").Json())
}

func UserInformationAPI(AccountId string) *gjson.Result {
	params := map[string]string{"expand": "introduction,bigAvatar,avatar,backgroundPic,fansNum,followNum,followyou,youfollow,verifyType,verifyInfo,avatarFrame,youblock,widgets,growup"}

	return VerifyAPI(request.Get("users/" + AccountId).AddAll(params).Json())
}

func UserWorksAPI(AccountId string) *gjson.Result {
	params := map[string]string{"expand": "typeName,sysTags,isbranch"}
	return VerifyAPI(request.Get("users/" + AccountId + "/novels").AddAll(params).Json())
}

func UserCommentAPI(AccountId string, page int) *gjson.Result {
	params := map[string]string{"expand": "novels,comics,albums,tags,sysTags,authorName", "size": "20", "page": fmt.Sprintf("%d", page)}
	return VerifyAPI(request.Get("users/" + AccountId + "/dynamics").AddAll(params).Json())
}

func LoginAccountAPI(username string, password string) string {
	var Cookie string
	response := request.Post("sessions").AddAll(map[string]string{"username": username, "password": password})
	if VerifyAPI(response.Json()) != nil {
		for _, cookie := range response.GetCookie() {
			Cookie += cookie.Name + "=" + cookie.Value + ";"
		}
		return Cookie
	}
	fmt.Println("login failed! error:", Template.Login.Status.Msg)
	return ""
}
