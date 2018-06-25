package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"

	"github.com/gotoronto/rnd-workshop/crawler/lists"
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
		body, err := get(urlToCrawl)
		if err != nil {
			log.Printf("Error crawling %s: %v\n", urlToCrawl, err)
			continue
		}

		for _, newURL := range parseLinks(body) {
			if ok, _ := crawler.list.Add(newURL); ok {
				log.Printf("Found new url: %s", newURL)
			}
		}

		err = body.Close()
		if err != nil {
			log.Printf("Failed to close body: %v\n", err)
		}

		urlToCrawl, moreUrls = crawler.list.Pop()
	}
}

func get(url string) (io.ReadCloser, error) {
	log.Printf("Visiting %s", url)
	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}

func parseLinks(r io.Reader) []string {
	tk := html.NewTokenizer(r)
	var links []string
	for {
		tt := tk.Next()

		switch tt {
		case html.ErrorToken:
			return links
		case html.StartTagToken:
			t := tk.Token()
			if t.Data != "a" {
				break
			}
			for _, a := range t.Attr {
				if a.Key != "href" {
					continue
				}

				parsed, err := url.ParseRequestURI(a.Val)
				if err != nil || !parsed.IsAbs() {
					continue
				}

				fmt.Println(parsed)
				links = append(links, parsed.String())
				break
			}
		}
	}
}
