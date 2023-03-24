package boluobaoapi

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/request"
	"github.com/tidwall/gjson"
)

func NovelCatalogueAPI(NovelID string) *gjson.Result {
	response := request.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Add("expand", "originNeedFireMoney").Json()
	return VerifyAPI(response)

}

func NovelContentAPI(chapter_id string) *gjson.Result {
	return VerifyAPI(request.Get("Chaps/"+chapter_id).Add("expand", "content").Json())
}
