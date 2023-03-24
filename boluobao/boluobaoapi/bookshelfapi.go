package boluobaoapi

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
)

func NovelBookShelfAPI() *gjson.Result {
	return VerifyAPI(request.Get("user/Pockets").Add("expand", "novels").Json())
}
