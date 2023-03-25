package boluobaoapi

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/request"
	"github.com/tidwall/gjson"
)

func NovelCatalogueAPI(NovelID string) *gjson.Result {
	response := request.Get(fmt.Sprintf("novels/%v/dirs", NovelID)).Data(map[string]string{"expand": "originNeedFireMoney"}).Json()
	return VerifyAPI(response)

}

func NovelContentAPI(chapter_id string) *gjson.Result {
	return VerifyAPI(request.Get("Chaps/" + chapter_id).Data(map[string]string{"expand": "content"}).Json())
}
