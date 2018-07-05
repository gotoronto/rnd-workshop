package crawler

import (
	"github.com/gotoronto/rnd-workshop/crawler/lists"
)

var List = lists.NewURLList()

func Add(url string) (bool, error) {
	return List.Add(url)
}

func Delete(url string) bool {
	return List.Delete(url)
}
