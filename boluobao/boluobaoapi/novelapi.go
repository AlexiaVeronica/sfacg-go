package boluobaoapi

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
	"strconv"
)

func NovelInformationAPI(NovelId string) *gjson.Result {
	expand := "chapterCount,bigBgBanner,bigNovelCover,typeName,intro,fav,ticket,pointCount,tags,sysTags,signlevel,discount,discountExpireDate,totalNeedFireMoney,rankinglist,originTotalNeedFireMoney,firstchapter,latestchapter,latestcommentdate,essaytag,auditCover,preOrderInfo,customTag,topic,unauditedCustomtag,homeFlag,isbranch"
	return VerifyAPI(request.Get(fmt.Sprintf("novels/%v", NovelId)).Add("expand", expand).Json())
}

func NovelPocketAPI(NovelId string) *gjson.Result {
	params := `{"novelId":` + NovelId + `,"categoryId":0}`
	return VerifyAPI(request.Post("pockets/-1/novels").AddString(params).Json())
}

func NovelFavAPI(NovelId string) *gjson.Result {
	return VerifyAPI(request.Post("novels/" + NovelId + "/favs").Json())
}

func NovelFeedAPI(page int, f bool) *gjson.Result {
	var filter string
	if f {
		filter = "hot"
	} else {
		filter = "latest"
	}
	expand := "novels,comics,albums,tags,sysTags,authorName,hasfollowed"
	return VerifyAPI(request.Get("user/feeds").AddAll(map[string]string{
		"filter": filter, "expand": expand, "page": strconv.Itoa(page), "size": "20"}).Json())
}
