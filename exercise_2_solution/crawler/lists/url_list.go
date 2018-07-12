package lists

import (
	"errors"
)

// URLList is a list of urls
type URLList struct {
	URLs []string
}

// NewURLList will create a new URLList with provided url
func NewURLList(url string) *URLList {
	urls := []string{}
	if url != "" {
		urls = []string{url}
	}
	return &URLList{URLs: urls}
}

// Add will add a url to the list if it is not already in the list
func (list *URLList) Add(url string) (bool, error) {
	if ok, _ := list.find(url); ok {
		return false, errors.New("url already exists")
	}
	list.URLs = append(list.URLs, url)
	return true, nil
}

// Find will check if the url exists already in the list
func (list *URLList) find(url string) (bool, int) {
	for i, u := range list.URLs {
		if u == url {
			return true, i
		}
	}
	return false, -1
}
