package bookscraper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"sync"
	"time"
)

type BookScraper struct {
	stores         map[string][]book
	ctx            context.Context
	allocationSize int
}

func New(allocationSize int) *BookScraper {
	return &BookScraper{
		stores:         make(map[string][]book),
		ctx:            context.Background(),
		allocationSize: allocationSize,
	}
}

func (bs *BookScraper) ScrapeAll(s string, scrapers []Scraper) {
	ctx, cancel := context.WithTimeout(bs.ctx, time.Second*3)
	defer cancel()

	respCh := make(chan Response)
	wg := &sync.WaitGroup{}
	wg.Add(len(scrapers))

	q := url.QueryEscape(s)

	for _, s := range scrapers {
		name := s.GetName()

		if bs.stores[name] != nil {
			panic("already presented in scraper")
		}

		// allocate x size of books for each store
		bs.stores[name] = make([]book, bs.allocationSize)

		// run each scraper in a separate goroutine
		go func(s Scraper) {
			s.Scrape(respCh, bs.stores[name], q)
			wg.Done()
		}(s)
	}

	// wait for finish here
	isDoneCh := make(chan int)
	go func() {
		wg.Wait()
		close(isDoneCh)
	}()

	for {
		select {
		// time's up
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				return
			}

		// this channel tracks errors that've been occurred in goroutines
		case resp := <-respCh:
			if resp.err != nil {
				bs.stores[resp.scraper.GetName()] = nil
				log.Printf("%s -> %s", resp.scraper.GetName(), resp.err)
				continue
			}

			panic(fmt.Sprintf("received `Response` without error from `%s`", resp.scraper.GetName()))

		// `wg.Wait()` is finished
		case <-isDoneCh:
			return
		}
	}
}

func (bs *BookScraper) AsJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(bs.stores)
}
