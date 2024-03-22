package desume

// PageNavParams contains information about navigating through the result pages.
type PageNavParams struct {
	Count   int    `json:"count"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"order_by"`
}
