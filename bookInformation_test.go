package sfacg

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao"
	"testing"
)

// BookDelite NovelDelite

func TestInformation(t *testing.T) {
	res := boluobao.NovelInit{BookId: "222494"}
	fmt.Println(res.NovelInfo().Get("novelName").String())
	fmt.Println("--------------------------------------------------")
}
