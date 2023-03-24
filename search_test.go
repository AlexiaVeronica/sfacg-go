package BoluobaoAPI_master

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(boluobao.NovelSearch("都市", 10))
	fmt.Println("--------------------------------------------------")
}
