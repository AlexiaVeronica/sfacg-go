package Template

type BookInfoData struct {
	AuthorID       int     `json:"authorId"`
	LastUpdateTime string  `json:"lastUpdateTime"`
	MarkCount      int     `json:"markCount"`
	NovelCover     string  `json:"novelCover"`
	BgBanner       string  `json:"bgBanner"`
	NovelID        int     `json:"novelId"`
	NovelName      string  `json:"novelName"`
	Point          float64 `json:"point"`
	IsFinish       bool    `json:"isFinish"`
	AuthorName     string  `json:"authorName"`
	CharCount      int     `json:"charCount"`
	ViewTimes      int     `json:"viewTimes"`
	TypeID         int     `json:"typeId"`
	AllowDown      bool    `json:"allowDown"`
	SignStatus     string  `json:"signStatus"`
	CategoryID     int     `json:"categoryId"`
	Expand         struct {
		TypeName      string        `json:"typeName"`
		Tags          []interface{} `json:"tags"`
		SysTags       []SysTagList  `json:"sysTags"`
		Ticket        int           `json:"ticket"`
		LatestChapter struct {
			Title   string `json:"title"`
			ChapID  int    `json:"chapId"`
			AddTime string `json:"addTime"`
		} `json:"latestChapter"`
	} `json:"expand"`
}

type BookInfo struct {
	Status Status       `json:"status"`
	Data   BookInfoData `json:"data"`
}
