package sfacg

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao"
	"testing"
)

func TestSearch(t *testing.T) {
	res := boluobao.NovelInit{BookId: "222494"}
	fmt.Println(res.NovelInfo().Get("novelName").String())
	//fmt.Println(boluobao.NovelSearch("武", 10))
	fmt.Println("--------------------------------------------------")
}
