package boluobaoapi

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
)

func NovelSpecialPushAPI() *gjson.Result {
	return VerifyAPI(request.Get("specialpush").Add("expand", "novels").Json())
}

func NovelActPushAPI() *gjson.Result {
	return VerifyAPI(request.Get("actpush").AddAll(map[string]string{"filter": "android", "page": "0", "size": "20"}).Json())
}

func NovelRecommendAPI(pushNames string) *gjson.Result {
	// pushNames: hotpush, newpush,
	return VerifyAPI(request.Get("novels/specialpushs").AddAll(
		map[string]string{"pushNames": pushNames, "page": "0", "size": "10", "expand": "sysTags,homeFlag"}).Json())
}
