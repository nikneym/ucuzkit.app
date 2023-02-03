package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Idefix struct{}

func (s *Idefix) GetName() string {
	return "idefix"
}

func (s *Idefix) Scrape(ch chan Response, books []book, query string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.idefix.com", "idefix.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Println("scraping idefix...")
	})

	i := 0
	c.OnHTML("#facetProducts .cart-product-box-view", func(h *colly.HTMLElement) {
		if i == len(books) {
			ch <- Response{scraper: s, err: nil}
			return
		}

		books[i] = book{
			Name:      h.ChildText(".product-info .box-title a"),
			Author:    h.ChildText(".pName a.who"),
			Publisher: h.ChildText(".manufacturerName"),
			Cover:     h.ChildAttr(".image-area img", "data-src"),
			// Idefix does not supply a `PriceOld` value so we'll just skip that
			PriceNew: h.ChildAttr("span#prices", "data-price"),
		}
		i++
	})

	// fail
	c.OnError(func(r *colly.Response, err error) {
		ch <- Response{scraper: s, err: err}
	})

	// success
	c.OnScraped(func(r *colly.Response) {
		ch <- Response{scraper: s, err: nil}
	})

	err := c.Visit(fmt.Sprintf("https://www.idefix.com/search?q=%s&redirect=search", query))
	if err != nil {
		ch <- Response{scraper: s, err: err}
		return
	}
}
