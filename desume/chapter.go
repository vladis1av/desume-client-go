package desume

// ChapterInfo contains short information about the manga chapter
type ChapterInfo struct {
	Vol  string  `json:"vol"`
	Ch   string  `json:"ch"`
	Name *string `json:"name"`
	Date string  `json:"date"`
}

// Chapter contains full information about the manga chapter
type Chapter struct {
	ID    int     `json:"id"`
	Vol   int     `json:"vol"`
	Ch    int     `json:"ch"`
	Title *string `json:"title"`
	Date  int64   `json:"date"`
	Check int     `json:"check"`
}

// Chapters contains information about the First, Last and Updated manga chapter
type Chapters struct {
	First   ChapterInfo `json:"first"`
	Last    ChapterInfo `json:"last"`
	Updated ChapterInfo `json:"updated"`
}

// ChapterPage contains information about the current chapter page
type ChapterPage struct {
	ID     int    `json:"id"`
	Page   int    `json:"page"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Img    string `json:"img"`
}

// ChaptersList contains information about the all chapters
type ChaptersList struct {
	Chapters
	Count int       `json:"count,omitempty"`
	List  []Chapter `json:"list,omitempty"`
}
