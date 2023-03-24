package boluobaoapi

import (
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
	"strconv"
)

func TagAPI(TagID string, page int) *gjson.Result {
	return VerifyAPI(request.Get("novels/0/sysTags/novels").AddAll(
		map[string]string{
			"systagids":      TagID,
			"isfree":         "both",
			"size":           "50",
			"charcountbegin": "0",
			"updatedays":     "-1",
			"expand":         "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,rankinglist,firstchapter,latestchapter,essaytag,auditCover,preOrderInfo,customTag,topic,isbranch",
			"sort":           "latest",
			"page":           strconv.Itoa(page),
			"charcountend":   "0",
			"isfinish":       "both",
		}).Json())
}

func SysTagAPI() *gjson.Result {
	return VerifyAPI(request.Get("novels/0/sysTags").Add("filter", "push").Json())
}
