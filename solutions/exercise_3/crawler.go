package crawler

import (
	"github.com/gotoronto/rnd-workshop/lib/scraper"
)

// Crawl will crawl a webpage, and crawl any links that are found in the body of
// that webpage
func Crawl(url string) {
	visited := 0
	requested := make(map[string]bool)
	s := scraper.New()

	go s.Scrape(url)
	requested[url] = true

	for visited != len(requested) {
		resp := <-s.Responses
		visited++

		for _, url := range resp {
			if found := requested[url]; !found {
				go s.Scrape(url)
				requested[url] = true
			}
		}
	}
}
