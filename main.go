package main

import (
	"os"

	"github.com/nikneym/ucuzkit.app/bookscraper"
)

func main() {
	bs := bookscraper.New(5)
	bs.ScrapeAll("1984", []bookscraper.Scraper{
		&bookscraper.Dr{},
		&bookscraper.Idefix{},
		&bookscraper.Ilknokta{},
		&bookscraper.Kitapyurdu{},
	})

	bs.AsJSON(os.Stdout)
}
