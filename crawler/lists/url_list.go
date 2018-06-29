package lists

import (
	"errors"
)

// URLList is a list of urls
type URLList struct {
	Urls []string
}

// NewURLList will create a new URLList with provided seeds
func NewURLList(seeds ...string) *URLList {
	if seeds == nil {
		seeds = make([]string, 0)
	}
	return &URLList{Urls: seeds}
}

// Add will add a url to the list if it is not already in the list
func (list *URLList) Add(url string) (bool, error) {
	if ok, _ := list.Check(url); ok {
		return false, errors.New("url already exists")
	}
	list.Urls = append(list.Urls, url)
	return true, nil
}

// Check will check if the url exists already in the list
func (list *URLList) Check(url string) (bool, int) {
	for i, u := range list.Urls {
		if u == url {
			return true, i
		}
	}
	return false, -1
}

// Delete will remove a url from the list
func (list *URLList) Delete(url string) bool {
	if ok, i := list.Check(url); ok {
		list.Urls = append(list.Urls[:i], list.Urls[i+1:]...)
		return true
	}
	return false
}
