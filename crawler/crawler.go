package crawler

import (
	"fmt"
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
	return &Crawler{
		list: lists.NewURLList(seeds...),
		urls: make(chan string, concurrencyLimit),
	}
}

func (crawler *Crawler) Crawl() {
	fmt.Println("Starting crawler")
	url, more := crawler.list.Pop()
	for more {
		crawler.urls <- url
		url, more = crawler.list.Pop()
	}

	for {
		select {
		case url := <-crawler.urls:
			go crawler.crawl(url)
		}
	}
}

func (crawler *Crawler) crawl(url string) {
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
