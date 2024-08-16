package desume

// Translator contains information about the translator
type Translator struct {
	ID   int64   `json:"id"`
	Name string  `json:"name"`
	Site *string `json:"site"`
}
