package boluobaoapi

import (
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/tidwall/gjson"
	"strconv"
)

func SearchAPI(keyword string, page int) *gjson.Result {
	params := map[string]string{"q": keyword, "page": strconv.Itoa(page), "size": "50"}
	return VerifyAPI(request.Get("search/novels/result").AddAll(params).Json())
}
