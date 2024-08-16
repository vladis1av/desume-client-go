package desume

// Genre contains information about the manga genre
type Genre struct {
	ID      int64  `json:"id"`
	Kind    string `json:"kind"`
	Text    string `json:"text"`
	Russian string `json:"russian"`
}
