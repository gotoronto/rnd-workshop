package urlset

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
