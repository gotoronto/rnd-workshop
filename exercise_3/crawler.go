package crawler

import (
	"github.com/gotoronto/rnd-workshop/lib/scraper"
)

func Crawl(url string) {
	// Implement this method, please use the s.Scrape method scrape a webpage
	s := scraper.New()

	// This will block because it is trying to write to s.Responses
	s.Scrape(url)
}
