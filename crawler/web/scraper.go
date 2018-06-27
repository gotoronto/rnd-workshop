package web

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html"
)

// Scrape will fetch an html page and scrape it for links. These links will be
// returned as a string array. If there was an error while requesting it will be
// returned as well
func Scrape(url string) ([]string, error) {
	client := http.Client{
		Timeout: time.Duration(4 * time.Second),
	}
	response, err := client.Get(url)
	if err != nil {
		return []string{}, err
	}
	defer response.Body.Close()
	return parseLinks(response.Body), nil
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

				links = append(links, parsed.String())
				break
			}
		}
	}
}
