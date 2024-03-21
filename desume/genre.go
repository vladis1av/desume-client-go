package desume

type Genre struct {
	ID      int    `json:"id"`
	Kind    string `json:"kind"`
	Text    string `json:"text"`
	Russian string `json:"russian"`
}
