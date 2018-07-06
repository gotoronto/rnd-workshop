package crawler

import (
	"log"
	"os"

	"github.com/gotoronto/rnd-workshop/exercise_3/crawler/lists"
	"github.com/gotoronto/rnd-workshop/exercise_3/crawler/web"
)

type Crawler struct {
	List *lists.URLList
}

func New(seeds ...string) *Crawler {
	return &Crawler{
		List: lists.NewURLList(seeds...),
	}
}

func (crawler *Crawler) Crawl(done chan os.Signal) {
	log.Println("Starting crawler")
	for i := 0; i < len(crawler.List.URLs); i++ {
		url := crawler.List.URLs[i]
		log.Printf("Visiting %s", url)
		links, err := web.Scrape(url)
		if err != nil {
			log.Printf("Error crawling %s: %v\n", url, err)
			continue
		}

		for _, newURL := range links {
			crawler.List.Add(newURL)
		}

		if len(done) > 0 {
			return
		}
	}
}
