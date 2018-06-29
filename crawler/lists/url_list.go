package lists

import (
	"errors"
	"sync"
)

// URLList is a list of urls
type URLList struct {
	urls  []string
	mutex *sync.Mutex
}

// NewURLList will create a new URLList with provided seeds
func NewURLList(seeds ...string) *URLList {
	if seeds == nil {
		seeds = make([]string, 0)
	}
	return &URLList{
		urls:  seeds,
		mutex: &sync.Mutex{},
	}
}

// Add will add a url to the list if it is not already in the list
func (list *URLList) Add(url string) (bool, error) {
	if ok, _ := list.Find(url); ok {
		return false, errors.New("url already exists")
	}
	list.mutex.Lock()
	list.urls = append(list.urls, url)
	list.mutex.Unlock()
	return true, nil
}

// Find will check if the url exists already in the list
func (list *URLList) Find(url string) (bool, int) {
	for i, u := range list.urls {
		if u == url {
			return true, i
		}
	}
	return false, -1
}

// Delete will remove a url from the list
func (list *URLList) Delete(url string) bool {
	if ok, i := list.Find(url); ok {
		list.mutex.Lock()
		list.urls = append(list.urls[:i], list.urls[i+1:]...)
		list.mutex.Unlock()
		return true
	}
	return false
}

func (list *URLList) Pop() (string, bool) {
	if len(list.urls) <= 0 {
		return "", false
	}
	list.mutex.Lock()
	url := list.urls[0]
	list.urls = list.urls[0+1:]
	list.mutex.Unlock()
	return url, true
}
