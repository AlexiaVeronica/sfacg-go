package boluobaoapi

import (
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/tidwall/gjson"
)

func NovelBookShelfAPI() *gjson.Result {
	//return VerifyAPI(request.Get("user/Pockets").Add("expand", "novels").Json())
	return VerifyAPI(request.Get("user/Pockets").Data(map[string]string{"expand": "novels"}).Json())
}
