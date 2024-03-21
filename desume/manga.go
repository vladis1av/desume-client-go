package desume

type MangaBase struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Russian       string  `json:"russian"`
	Image         Image   `json:"image"`
	URL           string  `json:"url"`
	Kind          string  `json:"kind"`
	Reading       string  `json:"reading"`
	Ongoing       int     `json:"ongoing"`
	Anons         int     `json:"anons"`
	Adult         int     `json:"adult"`
	AgeLimit      string  `json:"age_limit"`
	Status        string  `json:"status"`
	TransStatus   string  `json:"trans_status"`
	AiredOn       int64   `json:"aired_on"`
	ReleasedOn    int64   `json:"released_on"`
	Score         float64 `json:"score"`
	ScoreUsers    int     `json:"score_users"`
	Views         int     `json:"views"`
	Description   *string `json:"description"`
	Checked       int64   `json:"checked"`
	Updated       int64   `json:"updated"`
	Synonyms      *string `json:"synonyms"`
	ThreadID      *int    `json:"thread_id"`
	ShikimoriID   *int    `json:"shikimori_id"`
	MyAnimeListID *int    `json:"myanimelist_id"`
	MangaDexID    *string `json:"mangadex_id"`
}

type MangaChapter struct {
	MangaBase
	Genres       []Genre      `json:"genres"`
	Translator   []Translator `json:"translator"`
	ChaptersList `json:"chapters"`
	Pages        struct {
		Ch_curr Chapter       `json:"ch_curr"`
		Ch_prev Chapter       `json:"ch_prev"`
		Ch_next Chapter       `json:"ch_next"`
		List    []ChapterPage `json:"list"`
	} `json:"pages"`
}

type MangaInfo struct {
	MangaBase
	Genres       []Genre      `json:"genres"`
	Translator   []Translator `json:"translator"`
	ChaptersList `json:"chapters"`
}

type MangaFiltered struct {
	MangaBase
	Genres   string   `json:"genres"`
	Chapters Chapters `json:"chapters"`
}

type MangaChapterResponse struct {
	Response []MangaChapter `json:"response"`
}

type MangaInfoResponse struct {
	Response []MangaInfo `json:"response"`
}

type MangasFilteredResponse struct {
	PageNavParams PageNavParams `json:"pageNavParams"`
	MangaBase
	Genres   string   `json:"genres"`
	Chapters Chapters `json:"chapters"`
}
