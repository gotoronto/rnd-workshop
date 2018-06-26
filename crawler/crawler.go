package crawler

import (
	"log"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
	"github.com/gotoronto/rnd-workshop/crawler/web"
)

const concurrencyLimit = 20

type Crawler struct {
	list *lists.URLList
	urls chan string
}

func New(seeds ...string) *Crawler {
	urlChan := make(chan string, concurrencyLimit)
	for _, url := range seeds {
		urlChan <- url
	}

	return &Crawler{
		list: lists.NewURLList(),
		urls: urlChan,
	}
}

func (crawler *Crawler) Crawl() {
	log.Println("Starting crawler")
	for {
		select {
		case url := <-crawler.urls:
			go crawler.crawl(url)
		}
	}
}

func (crawler *Crawler) crawl(url string) {
	log.Printf("Visiting %s\n", url)
	links, err := web.Scrape(url)
	if err != nil {
		log.Printf("Error crawling %s: %v\n", url, err)
		return
	}

	for _, newURL := range links {
		if ok, _ := crawler.list.Add(newURL); ok {
			crawler.urls <- newURL
		}
	}
}
