package boluobao

import (
	"fmt"
	"github.com/VeronicaAlexia/sfacg-go/boluobao/boluobaoapi"
	"github.com/VeronicaAlexia/sfacg-go/pkg/file"
	"github.com/VeronicaAlexia/sfacg-go/pkg/threading"
	"github.com/tidwall/gjson"
	"strings"
)

type NovelInit struct {
	BookId string
	Book   BookInfo
}
type BookInfo struct {
	BookId         string
	BookName       string
	Author         string
	Cover          string
	LastUpdateTime string
	MarkCount      string
	IsFinish       bool
	CharCount      string
	ViewTimes      string
	AuthorId       string
	SignStatus     string
}

func (book *NovelInit) NovelInfo() *NovelInit {
	if result := boluobaoapi.NovelInformationAPI(book.BookId); result != nil {
		book.Book = BookInfo{
			BookId:         result.Get("data.novelId").String(),
			BookName:       result.Get("data.novelName").String(),
			Cover:          result.Get("data.novelCover").String(),
			LastUpdateTime: result.Get("data.lastUpdateTime").String(),
			MarkCount:      result.Get("data.markCount").String(),
			IsFinish:       result.Get("data.isFinish").Bool(),
			CharCount:      result.Get("data.charCount").String(),
			ViewTimes:      result.Get("data.viewTimes").String(),
			AuthorId:       result.Get("data.authorId").String(),
			SignStatus:     result.Get("data.signStatus").String(),
		}
	} else {
		fmt.Println("获取小说信息失败,请检查小说ID是否正确:", book.BookId)
	}
	return book
}

func (book *NovelInit) NovelCatalogue() []string {
	var chapter_id_list []string
	if response := boluobaoapi.NovelCatalogueAPI(book.BookId); response != nil {
		for _, volume := range response.Get("data.volumeList").Array() {
			for _, chapter := range volume.Get("chapterList").Array() {
				if chapter.Get("originNeedFireMoney").Int() == 0 {
					chapter_id_list = append(chapter_id_list, chapter.Get("chapId").String())
				}
			}
		}
		return chapter_id_list
	}
	return nil
}

func (book *NovelInit) NovelContent(chapter_id string) string {
	var contentText string
	if response := boluobaoapi.NovelContentAPI(chapter_id); response != nil {
		for _, content := range strings.Split(response.Get("data.expand.content").String(), "\n") {
			if content != "" {
				contentText += "　　" + content + "\n"
			}
		}
		return response.Get("data.title").String() + "\n\n" + contentText

	}
	return ""
}

func (book *NovelInit) NovelDownload(path string) {
	fmt.Println("正在获取小说信息...")
	for _, chapter_id := range book.NovelCatalogue() {
		content := book.NovelContent(chapter_id)
		if content != "" {
			file.AppendToFile(path+book.Book.BookName+".txt", content)
		}

	}
}

func NovelSearch(keyword string, lastPage int) []gjson.Result {
	var BookInfoList []gjson.Result
	Thread := threading.InitThreading(32)
	for i := 0; i < lastPage; i++ {
		Thread.Add()
		go func(page int) {
			defer Thread.Done()
			response := boluobaoapi.SearchAPI(keyword, page)
			if response != nil {
				for _, data := range response.Get("data.novels").Array() {
					BookInfoList = append(BookInfoList, data)
				}
			}
		}(i)
	}
	Thread.Wait() // wait for all goroutines to finish
	return BookInfoList
}
