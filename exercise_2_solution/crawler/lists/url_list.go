package lists

import (
	"errors"
)

// URLList is a list of urls
type URLList struct {
	URLs []string
}

// NewURLList will create a new URLList with provided seeds
func NewURLList(seeds ...string) *URLList {
	if seeds == nil {
		seeds = make([]string, 0)
	}
	return &URLList{URLs: seeds}
}

// Add will add a url to the list if it is not already in the list
func (list *URLList) Add(url string) (bool, error) {
	if ok, _ := list.Find(url); ok {
		return false, errors.New("url already exists")
	}
	list.URLs = append(list.URLs, url)
	return true, nil
}

// Find will check if the url exists already in the list
func (list *URLList) Find(url string) (bool, int) {
	for i, u := range list.URLs {
		if u == url {
			return true, i
		}
	}
	return false, -1
}

// Delete will remove a url from the list
func (list *URLList) Delete(url string) bool {
	if ok, i := list.Find(url); ok {
		list.URLs = append(list.URLs[:i], list.URLs[i+1:]...)
		return true
	}
	return false
}
