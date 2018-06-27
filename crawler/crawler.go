package crawler

import (
	"log"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
	"github.com/gotoronto/rnd-workshop/crawler/web"
)

type Crawler struct {
	list    *lists.URLList
	visited *lists.URLList
}

func New(seeds ...string) *Crawler {
	return &Crawler{
		list:    lists.NewURLList(seeds...),
		visited: lists.NewURLList(),
	}
}

func (crawler *Crawler) Crawl(done chan int) {
	log.Println("Starting crawler")
	urlToCrawl, moreUrls := crawler.list.Pop()
	for moreUrls {
		log.Printf("Visiting %s", urlToCrawl)
		links, err := web.Scrape(urlToCrawl)
		if err != nil {
			log.Printf("Error crawling %s: %v\n", urlToCrawl, err)
			continue
		}

		for _, newURL := range links {
			if ok, err := crawler.visited.Add(newURL); ok && err == nil {
				crawler.list.Add(newURL)
			}
		}
		urlToCrawl, moreUrls = crawler.list.Pop()
	}
}
