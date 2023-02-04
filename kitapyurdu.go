package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Kitapyurdu struct{}

func (s *Kitapyurdu) GetName() string {
	return "kitapyurdu"
}

func (s *Kitapyurdu) Scrape(ch chan Response, books []book, query string) {
	c := colly.NewCollector(
		colly.AllowedDomains("kitapyurdu.com", "www.kitapyurdu.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Println("scraping kitapyurdu...")
	})

	i := 0
	c.OnHTML("ul.product-grid .product-cr, #product-table .product-cr", func(h *colly.HTMLElement) {
		if i == len(books) {
			return
		}

		books[i] = book{
			Name:      h.ChildText(".name"),
			Author:    h.ChildText(".author span a span"),
			Publisher: h.ChildText(".publisher"),
			Cover:     h.ChildAttr(".image .cover a img", "src"),
			PriceOld:  h.ChildText(".price .price-old span.value"),
			PriceNew:  h.ChildText(".price .price-new span.value"),
		}
		i++
	})

	// fail
	c.OnError(func(r *colly.Response, err error) {
		ch <- Response{scraper: s, err: err}
	})

	c.Visit(fmt.Sprintf("https://www.kitapyurdu.com/index.php?route=product/search&filter_name=%s&fuzzy=0&filter_product_type=1&filter_in_shelf=1", query))
}
