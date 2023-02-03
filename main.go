package main

import "fmt"

func main() {
	bs := NewBookScraper(5)
	bs.ScrapeAll([]Scraper{
		&Dr{},
		&Idefix{},
		&Ilknokta{},
		&Kitapyurdu{},
	})

	fmt.Println(bs.stores)
}
