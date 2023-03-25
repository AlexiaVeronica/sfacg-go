package boluobaoapi

import (
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/tidwall/gjson"
)

func NovelSpecialPushAPI() *gjson.Result {
	return VerifyAPI(request.Get("specialpush").Data(map[string]string{"expand": "novels"}).Json())
}

func NovelActPushAPI() *gjson.Result {
	return VerifyAPI(request.Get("actpush").Data(map[string]string{"filter": "android", "page": "0", "size": "20"}).Json())
}

func NovelRecommendAPI(pushNames string) *gjson.Result {
	// pushNames: hotpush, newpush,
	return VerifyAPI(request.Get("novels/specialpushs").Data(
		map[string]string{"pushNames": pushNames, "page": "0", "size": "10", "expand": "sysTags,homeFlag"}).Json())
}
