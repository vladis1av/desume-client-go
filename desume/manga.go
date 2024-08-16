package desume

import "encoding/json"

// MangaBase contains basic information about manga.
type MangaBase struct {
	ID            int64   `json:"id"`
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
	ThreadID      *int64  `json:"thread_id,omitempty"`
	ReleasedOn    int64   `json:"released_on"`
	ScoreUsers    int     `json:"score_users"`
	Description   *string `json:"description"`
	MangaDexID    *string `json:"mangadex_id"`
	TransStatus   string  `json:"trans_status"`
	ShikimoriID   *int64  `json:"shikimori_id,omitempty"`
	MyAnimeListID *int64  `json:"myanimelist_id,omitempty"`
}

// MaybeChapter is a structure that can contain either a value of type Chapter or // the special value "-1" to indicate no value.
type MaybeChapter struct {
	Value  *Chapter
	Exists bool
}

// UnmarshalJSON implements the json.Unmarshaler interface for the MaybeChapter type.
// // This method is used to deserialize JSON data into the Maybe Chapter structure.
// If the JSON data is a string "-1", then Value is set to nil and Exists to false.
// Otherwise, the JSON data is deserialized into the Chapter structure, and the reference to this structure
// is stored in Value, and Exists is set to true.
func (mc *MaybeChapter) UnmarshalJSON(data []byte) error {
	if string(data) == "-1" {
		mc.Exists = false
		mc.Value = nil
		return nil
	}

	var ch Chapter
	if err := json.Unmarshal(data, &ch); err != nil {
		return err
	}
	mc.Value = &ch
	mc.Exists = true
	return nil
}

// MarshalJSON implements the json.Marshaler interface for the MaybeChapter type. // This method is used to serialize the MaybeChapter structure into JSON data. // If Exists is false, then the method returns the string "-1". // Otherwise, the method serializes the value contained in Value into JSON data.
func (mc MaybeChapter) MarshalJSON() ([]byte, error) {
	if !mc.Exists {
		return []byte("-1"), nil
	}
	return json.Marshal(mc.Value)
}

// MangaChapterPages provides information about the current, previous, and next manga chapters,
// as well as a list of pages for the current chapter.
type MangaChapterPages struct {
	Ch_curr Chapter       `json:"ch_curr,omitempty"`
	Ch_prev MaybeChapter  `json:"ch_prev,omitempty"`
	Ch_next MaybeChapter  `json:"ch_next,omitempty"`
	List    []ChapterPage `json:"list,omitempty"`
}

// MangaChapter contains information about the manga and chapters.
type MangaChapter struct {
	MangaBase
	Genres       []Genre            `json:"genres"`
	Translator   []Translator       `json:"translator"`
	Pages        *MangaChapterPages `json:"pages,omitempty"`
	ChaptersList `json:"chapters"`
}

// MangaInfo contains full information about the manga.
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
