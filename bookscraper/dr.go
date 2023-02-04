package bookscraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Dr struct{}

func (s *Dr) GetName() string {
	return "dr"
}

func (s *Dr) Scrape(ch chan Response, books []book, query string) {
	c := colly.NewCollector(
		colly.AllowedDomains("dr.com.tr", "www.dr.com.tr"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Println("scraping dr...")
	})

	i := 0
	c.OnHTML(".prd-main-wrapper", func(h *colly.HTMLElement) {
		if i == len(books) {
			return
		}

		books[i] = book{
			Name:      h.ChildText("a.prd-name"),
			Author:    h.ChildText(".pName a.who"),
			Cover:     h.ChildAttr(".product-img img", "data-src"),
			Publisher: h.ChildAttr(".prd-info-wrapper a.prd-publisher", "title"),
			// Dr does not supply a `PriceOld` value so we'll just skip that
			PriceNew: h.ChildAttr(".prd-price-wrapper .prd-prices .prd-price", "data-price"),
		}
		i++
	})

	// fail
	c.OnError(func(r *colly.Response, err error) {
		ch <- Response{scraper: s, err: err}
	})

	c.Visit(fmt.Sprintf("https://www.dr.com.tr/search?q=%s&redirect=search", query))
}
