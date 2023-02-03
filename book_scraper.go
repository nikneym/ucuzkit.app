package main

import (
	"context"
	"time"
)

type BookScraper struct {
	stores         map[string][]book
	ctx            context.Context
	allocationSize int
}

func NewBookScraper(allocationSize int) *BookScraper {
	return &BookScraper{
		stores:         make(map[string][]book),
		ctx:            context.Background(),
		allocationSize: allocationSize,
	}
}

func (bs *BookScraper) ScrapeAll(scrapers []Scraper) {
	ctx, cancel := context.WithTimeout(bs.ctx, time.Second*1)
	defer cancel()

	ch := make(chan Response)

	for _, s := range scrapers {
		name := s.GetName()

		if bs.stores[name] != nil {
			panic("already presented in scraper")
		}

		// allocate x size of books for each store
		bs.stores[name] = make([]book, bs.allocationSize)

		// run each scraper in a separate goroutine
		go s.Scrape(ch, bs.stores[name], "1984")
	}

	//tasksDone := 0
	//length := len(scrapers)
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return
			}

			//case resp := <-ch:
			//	if resp.err != nil {
			//		log.Fatal(resp.err)
			//		return
			//	}
			//
			//	if tasksDone > length {
			//		return
			//	}
			//
			//	tasksDone++
		}
	}
}
