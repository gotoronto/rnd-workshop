package crawler

import (
	"github.com/gotoronto/rnd-workshop/crawler/lists"
)

var list = lists.NewURLList()

func Add(url string) (bool, error) {
	return list.Add(url)
}

func Delete(url string) bool {
	return list.Delete(url)
}
