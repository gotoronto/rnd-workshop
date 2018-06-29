package crawler

import (
	"log"
	"os"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
	"github.com/gotoronto/rnd-workshop/crawler/web"
)

const concurrencyLimit = 20

type Crawler struct {
	urls chan string
	List *lists.URLList
}

func New(seeds ...string) *Crawler {
	urlChan := make(chan string, concurrencyLimit)
	for _, url := range seeds {
		urlChan <- url
	}

	return &Crawler{
		urls: urlChan,
		List: lists.NewURLList(),
	}
}

func (crawler *Crawler) Crawl(done chan os.Signal) {
	log.Println("Starting crawler")
	for {
		select {
		case url := <-crawler.urls:
			if found, _ := crawler.List.Find(url); !found {
				go crawler.crawl(url)
			}
		case <-done:
			return
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

	crawler.List.Add(url)

	for _, newURL := range links {
		crawler.urls <- newURL
	}
}
