package desume

// MangaBase contains basic information about manga.
type MangaBase struct {
	ID            int     `json:"id"`
	URL           string  `json:"url"`
	Name          string  `json:"name"`
	Kind          string  `json:"kind"`
	Image         Image   `json:"image"`
	Anons         int     `json:"anons"`
	Adult         int     `json:"adult"`
	Score         float64 `json:"score"`
	Views         int     `json:"views"`
	Status        string  `json:"status"`
	Russian       string  `json:"russian"`
	Reading       string  `json:"reading"`
	Checked       int64   `json:"checked"`
	Ongoing       int     `json:"ongoing"`
	Updated       int64   `json:"updated"`
	AiredOn       int64   `json:"aired_on"`
	Synonyms      *string `json:"synonyms"`
	AgeLimit      string  `json:"age_limit"`
	ThreadID      *int    `json:"thread_id,omitempty"`
	ReleasedOn    int64   `json:"released_on"`
	ScoreUsers    int     `json:"score_users"`
	Description   *string `json:"description"`
	MangaDexID    *string `json:"mangadex_id"`
	TransStatus   string  `json:"trans_status"`
	ShikimoriID   *int    `json:"shikimori_id,omitempty"`
	MyAnimeListID *int    `json:"myanimelist_id,omitempty"`
}

// MangaChapterPages contains information about the pages of the manga chapter.
type MangaChapterPages struct {
	Ch_curr Chapter       `json:"ch_curr,omitempty"`
	Ch_prev Chapter       `json:"ch_prev,omitempty"`
	Ch_next Chapter       `json:"ch_next,omitempty"`
	List    []ChapterPage `json:"list,omitempty"`
}

// MangaChapterPages contains information about the manga and chapters.
type MangaChapter struct {
	MangaBase
	Genres       []Genre           `json:"genres"`
	Translator   []Translator      `json:"translator"`
	Pages        MangaChapterPages `json:"pages,omitempty"`
	ChaptersList `json:"chapters"`
}

// MangaChapterPages contains full information about the manga.
type MangaInfo struct {
	MangaBase
	Genres       []Genre      `json:"genres"`
	Translator   []Translator `json:"translator"`
	ChaptersList `json:"chapters"`
}

// MangaFiltered contains information about the manga, filtered
// according to the specified params.
type MangaFiltered struct {
	MangaBase
	Genres   string   `json:"genres"`
	Chapters Chapters `json:"chapters"`
}

// Manga Chapter Response is an API response containing
// information about one chapter of the manga. Including the manga itself.
type MangaChapterResponse struct {
	Response MangaChapter `json:"response"`
}

// MangaInfoResponse presents an API response containing
// detailed information about a single manga.
type MangaInfoResponse struct {
	Response MangaInfo `json:"response"`
}

// MangasFilteredResponse represents an API response containing a filtered list of manga.
type MangasFilteredResponse struct {
	PageNavParams PageNavParams   `json:"pageNavParams"`
	Response      []MangaFiltered `json:"response"`
}
