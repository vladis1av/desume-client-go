package desume

type ChapterInfo struct {
	Vol  string  `json:"vol"`
	Ch   string  `json:"ch"`
	Name *string `json:"name"`
	Date string  `json:"date"`
}

type Chapter struct {
	ID    int     `json:"id"`
	Vol   int     `json:"vol"`
	Ch    int     `json:"ch"`
	Title *string `json:"title"`
	Date  int64   `json:"date"`
	Check int     `json:"check"`
}

type Chapters struct {
	First   ChapterInfo `json:"first"`
	Last    ChapterInfo `json:"last"`
	Updated ChapterInfo `json:"updated"`
}

type ChapterPage struct {
	ID     int    `json:"id"`
	Page   int    `json:"page"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Img    string `json:"img"`
}

type ChaptersList struct {
	Chapters
	Count int       `json:"count,omitempty"`
	List  []Chapter `json:"list,omitempty"`
}
