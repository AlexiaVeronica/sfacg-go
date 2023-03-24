package boluobaoapi

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
	"strconv"
)

func SearchAPI(keyword string, page int) *gjson.Result {
	return VerifyAPI(request.Get("search/novels/result").AddAll(
		map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "50"}).Json())
}
