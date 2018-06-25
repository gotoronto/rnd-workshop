package crawler

import (
	"fmt"
	"log"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
	"github.com/gotoronto/rnd-workshop/crawler/web"
)

type Crawler struct {
	list *lists.URLList
}

func New(seeds ...string) *Crawler {
	return &Crawler{
		list: lists.NewURLList(seeds...),
	}
}

func (crawler *Crawler) Crawl() {
	fmt.Println("Starting crawler")
	urlToCrawl, moreUrls := crawler.list.Pop()
	for moreUrls {
		links, err := web.Scrape(urlToCrawl)
		if err != nil {
			log.Printf("Error crawling %s: %v\n", urlToCrawl, err)
			continue
		}

		for _, newURL := range links {
			if ok, _ := crawler.list.Add(newURL); ok {
				log.Printf("Found new url: %s", newURL)
			}
		}
		urlToCrawl, moreUrls = crawler.list.Pop()
	}
}
