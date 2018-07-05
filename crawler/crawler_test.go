package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// To run the test
// dep ensure			// to install dependencies
// go test ./...

// Resources:
// https://gobyexample.com/
// https://golang.org/doc/ (use the search bar for doc)
// Google and Stack Overflow are your friends :)

// SOLUTION:
// func Add(url string) {
//	URLs = append(URLs, url)
// }
func TestAddANewURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// SOLUTION:
// func Add(url string) {
//	for _, u := range URLs {
//		if u == url {
//			return
//		}
//	}

//	URLs = append(URLs, url)
// }
func TestAddOnlyKeepsUniqueURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// doc on deleting an element from an array: https://github.com/golang/go/wiki/SliceTricks#delete

// SOLUTION:
// func Delete(url string) {
//	var i int
//	var u string
//	for i, u = range URLs {
//		if u == url {
//			break
//		}
//	}
//	URLs = append(URLs[:i], URLs[i+1:]...)
// }
func TestDeleteURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
}

// SOLUTION:
// func Delete(url string) {
//	var i int
//	var u string
//	for _, u = range URLs {
//		if u == url {
//			break
//		}
//		i++
//	}

//	if i >= len(URLs) {
//		return
//	}

//	URLs = append(URLs[:i], URLs[i+1:]...)
// }
func TestDeleteDoesNothingIfURLNotExist(t *testing.T) {
	Add("http://www.amazon.ca")
	Delete("http://www.google.ca")

	assert.Equal(t, []string{"http://www.amazon.ca"}, URLs)
}

// SOLUTION:
// func Delete(url string) bool {
//	var i int
//	var u string
//	for _, u = range URLs {
//		if u == url {
//			break
//		}
//		i++
//	}

//	if i >= len(URLs) {
//		return false
//	}

//	URLs = append(URLs[:i], URLs[i+1:]...)
//	return true
// }
func TestDeleteReturnsTrueIfDeleteIsPerformed(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	assert.True(t, deleted)
}

// solution same as previous
func TestDeleteReturnsFalseIfDeleteIsNotPerformed(t *testing.T) {
	URLs = []string{}
	deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	assert.False(t, deleted)
}

// SOLUTION:
// import "errors"

// func Add(url string) (bool, error) {
//	for _, u := range URLs {
//		if u == url {
//			return false, errors.New("url already exists.")
//		}
//	}

//	URLs = append(URLs, url)
//	return true, nil
// }
func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
	URLs = []string{"http://www.google.ca"}
	added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.False(t, added)
	assert.Equal(t, err.Error(), "url already exists.")
}

// Solution same as previous
func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
	URLs = []string{}
	added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.True(t, added)
	assert.Nil(t, err)
}
