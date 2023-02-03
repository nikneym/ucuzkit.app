package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Ilknokta struct{}

func (s *Ilknokta) GetName() string {
	return "ilknokta"
}

func (s *Ilknokta) Scrape(ch chan Response, books []book, query string) {
	c := colly.NewCollector(
		colly.AllowedDomains("ilknokta.com", "www.ilknokta.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.Println("scraping ilknokta...")
	})

	i := 0
	c.OnHTML("ul li.items_col .home_item_prd", func(h *colly.HTMLElement) {
		if i == len(books) {
			ch <- Response{scraper: s, err: nil}
			return
		}

		// TODO: parse book covers too
		books[i] = book{
			Name:      h.ChildText(".prd_info .name a"),
			Author:    h.ChildText(".prd_info .writer a"),
			Publisher: h.ChildText(".prd_info .publisher a"),
			PriceOld:  h.ChildAttr(".prd_info .price_box .price_box_wrapper span[class=\"price price_list convert_cur\"]", "data-price"),
			PriceNew:  h.ChildAttr(".prd_info .price_box .price_box_wrapper span[class=\"price price_sale convert_cur\"]", "data-price"),
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

	err := c.Visit(fmt.Sprintf("https://www.ilknokta.com/index.php?p=Products&q_field_active=0&q=%s&search=&q_field=", query))
	if err != nil {
		ch <- Response{scraper: s, err: err}
		return
	}
}
