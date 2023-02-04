package bookscraper

type Scraper interface {
	Scrape(ch chan Response, books []book, query string)
	GetName() string
}

type Response struct {
	scraper Scraper
	err     error
}
