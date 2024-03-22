package desume

// Image contains information about the manga covers
type Image struct {
	Original string `json:"original"`
	Preview  string `json:"preview"`
	X225     string `json:"x225"`
	X120     string `json:"x120"`
	X48      string `json:"x48"`
	X32      string `json:"x32"`
}
