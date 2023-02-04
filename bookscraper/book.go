package bookscraper

// FIXME: should price be stored as a string or as a decimal value?
type book struct {
	Name      string `json:"name,omitempty"`
	Author    string `json:"author,omitempty"`
	Publisher string `json:"publisher,omitempty"`
	Cover     string `json:"cover,omitempty"`
	PriceOld  string `json:"priceOld,omitempty"`
	PriceNew  string `json:"priceNew,omitempty"`
}
