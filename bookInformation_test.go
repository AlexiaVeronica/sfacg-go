package sfacg

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao"
	"testing"
)

func TestInformation(t *testing.T) {
	res := boluobao.NovelInit{BookId: "222494"}
	fmt.Println(res.NovelInfo().Book.LastUpdateTime)
	fmt.Println("--------------------------------------------------")
}
