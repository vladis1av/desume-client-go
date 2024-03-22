package desume

// Translator contains information about the translator
type Translator struct {
	ID   int     `json:"id"`
	Name string  `json:"name"`
	Site *string `json:"site"`
}
