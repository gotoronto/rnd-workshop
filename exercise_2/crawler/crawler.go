package crawler

import "errors"

var URLs []string

func Add(url string) (bool, error) {
	for _, u := range URLs {
		if u == url {
			return false, errors.New("url already exists.")
		}
	}

	URLs = append(URLs, url)
	return true, nil
}

func Delete(url string) bool {
	var i int
	var u string
	for _, u = range URLs {
		if u == url {
			break
		}
		i++
	}

	if i >= len(URLs) {
		return false
	}

	URLs = append(URLs[:i], URLs[i+1:]...)
	return true
}
