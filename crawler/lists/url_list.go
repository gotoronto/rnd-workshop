package lists

import (
	"errors"
)

// URLList is a list of urls
type URLList struct {
	urls []string
}

// NewURLList will create a new URLList with provided seeds
func NewURLList(seeds ...string) *URLList {
	if seeds == nil {
		seeds = make([]string, 0)
	}
	return &URLList{urls: seeds}
}

// Add will add a url to the list if it is not already in the list
func (list *URLList) Add(url string) (bool, error) {
	if ok, _ := list.Check(url); ok {
		return false, errors.New("url already exists")
	}
	list.urls = append(list.urls, url)
	return true, nil
}

// Check will check if the url exists already in the list
func (list *URLList) Check(url string) (bool, int) {
	for i, u := range list.urls {
		if u == url {
			return true, i
		}
	}
	return false, -1
}

// Delete will remove a url from the list
func (list *URLList) Delete(url string) bool {
	if ok, i := list.Check(url); ok {
		list.urls = append(list.urls[:i], list.urls[i+1:]...)
		return true
	}
	return false
}

func (list *URLList) Pop() (string, bool) {
	if len(list.urls) <= 0 {
		return "", false
	}
	url := list.urls[0]
	list.urls = list.urls[0+1:]
	return url, true
}
