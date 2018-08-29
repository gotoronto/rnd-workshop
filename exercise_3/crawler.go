package crawler

import (
	"net/http"

	"golang.org/x/net/html"
)

func Crawl(url string) {
	// Implement this method, please use the Scrape method below to scrape a webpage
	// at the url
}

// Scrape will perform a GET request to fetch an html page and scrape it for links.
// These links will be returned as a string array. If there was an error while
// requesting it will be returned as well
func Scrape(uri string, responses chan []string) {
	links := []string{}
	response, err := http.Get(uri) // request the uri
	if err != nil {
		responses <- links
		return
	}
	defer response.Body.Close()
	tokenizer := html.NewTokenizer(response.Body) // new html tokenizer to find tags
	for {
		tokenType := tokenizer.Next()     // get the next token type
		if tokenType == html.ErrorToken { // this indicates invalid html or complete scanning
			break
		}
		token := tokenizer.Token()                                             // get the html token we are looking at
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" { // if this is a link tag
			for _, a := range token.Attr { // look through the tags attributes
				if a.Key == "href" { // if we have the link value append it to our links
					links = append(links, a.Val)
				}
			}
		}
	}
	responses <- links
}
