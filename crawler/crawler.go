package crawler

import (
	"log"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
	"github.com/gotoronto/rnd-workshop/crawler/web"
)

type Crawler struct {
	List *lists.URLList
}

func New(seeds ...string) *Crawler {
	return &Crawler{
		List: lists.NewURLList(seeds...),
	}
}

func (crawler *Crawler) Crawl(done chan int) {
	log.Println("Starting crawler")
	for _, url := range crawler.List.URLs {
		log.Printf("Visiting %s", url)
		links, err := web.Scrape(url)
		if err != nil {
			log.Printf("Error crawling %s: %v\n", url, err)
			continue
		}

		for _, newURL := range links {
			crawler.List.Add(newURL)
		}
	}
}
