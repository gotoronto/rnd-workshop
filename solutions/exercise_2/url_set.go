package urlset

import (
	"errors"
)

// URLSet is a list of urls
type URLSet struct {
	URLs []string
}

// NewURLSet will create a new URLSet with provided seeds
func NewURLSet() *URLSet {
	return &URLSet{URLs: make([]string, 0)}
}

// Add will add a url to the list if it is not already in the list
func (list *URLSet) Add(url string) (bool, error) {
	for _, u := range list.URLs {
		if u == url {
			return false, errors.New("url already exists")
		}
	}

	list.URLs = append(list.URLs, url)
	return true, nil
}
