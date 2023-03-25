package boluobaoapi

import (
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/tidwall/gjson"
	"strconv"
)

func SearchAPI(keyword string, page int) *gjson.Result {
	params := map[string]string{
		"q":          keyword,
		"expand":     "novels,comics,albums,chatnovelstags,typeName,authorName,intro,latestchaptitle,latestchapintro,tags,sysTags",
		"sort":       "hot",
		"page":       strconv.Itoa(page),
		"size":       "12",
		"systagids":  "",
		"isFinish":   "-1",
		"updateDays": "-1",
	}
	return VerifyAPI(request.Get("search/novels/result/new").Data(params).Json())

}
