package sfacg

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao"
	"testing"
)

func TestSearch(t *testing.T) {
	fmt.Println(boluobao.NovelSearch("血姬", 10))
	fmt.Println("--------------------------------------------------")
}
