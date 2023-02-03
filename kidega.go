// TODO: rewrite in Scraper fashion
package main

//func ScrapeKidegaBooks(ch chan error, books []book, name string) {
//	defer close(ch)
//
//	c := colly.NewCollector(
//		colly.AllowedDomains("kidega.com", "www.kidega.com"),
//	)
//
//	c.OnRequest(func(r *colly.Request) {
//		log.Println("scraping kidega...")
//	})
//
//	i := 0
//	c.OnHTML("ul.product-list-grid li.prd .prd-inner", func(h *colly.HTMLElement) {
//		if i == len(books) {
//			ch <- nil
//			return
//		}
//
//		books[i] = book{
//			Name:      h.ChildText(".prd-body .authorArea .manufacturer a span.manufacturer-name"),
//			Publisher: h.ChildText(".prd-body .publisherArea .distributor a span.distributor-name"),
//		}
//
//		// publisher
//		fmt.Println(h.ChildText(".prd-footer .prd-price .first div[class=\"first-inner priceBar sale\"] .firstPrice"))
//
//		i++
//	})
//
//	// fail
//	c.OnError(func(r *colly.Response, err error) {
//		ch <- err
//	})
//
//	// success
//	c.OnScraped(func(r *colly.Response) {
//		ch <- nil
//	})
//
//	c.Request(
//		"GET",
//		fmt.Sprintf("https://kidega.com/arama/%s/", name),
//		nil,
//		nil,
//		http.Header{"User-Agent": []string{"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"}},
//	)
//}
//
