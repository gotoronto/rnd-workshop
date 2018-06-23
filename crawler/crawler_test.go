package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: expand the description of the test in a comment
// TODO: add "SOLUTION:" before each solution
// TODO: add a blurb about cd into the repo and how to run test
// TODO: add go dep

// func Add(url string) {
// 	URLs = append(URLs, url)
// }
func TestAddANewURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// func Add(url string) {
// 	for _, u := range URLs {
// 		if u == url {
// 			return
// 		}
// 	}

// 	URLs = append(URLs, url)
// }
func TestAddOnlyKeepsUniqueURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// doc on deleting an element from an array: https://github.com/golang/go/wiki/SliceTricks#delete

// func Delete(url string) {
// 	var i int
// 	var u string
// 	for i, u = range URLs {
// 		if u == url {
// 			break
// 		}
// 	}
// 	URLs = append(URLs[:i], URLs[i+1:]...)
// }
func TestDeleteURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
}

// func Delete(url string) {
// 	var i int
// 	var u string
// 	for _, u = range URLs {
// 		if u == url {
// 			break
// 		}
// 		i++
// 	}

// 	if i >= len(URLs) {
// 		return
// 	}

// 	URLs = append(URLs[:i], URLs[i+1:]...)
// }
func TestDeleteDoesNothingIfURLNotExist(t *testing.T) {
	Add("http://www.amazon.ca")
	Delete("http://www.google.ca")

	assert.Equal(t, []string{"http://www.amazon.ca"}, URLs)
}

// func Delete(url string) bool {
// 	var i int
// 	var u string
// 	for _, u = range URLs {
// 		if u == url {
// 			break
// 		}
// 		i++
// 	}

// 	if i >= len(URLs) {
// 		return false
// 	}

// 	URLs = append(URLs[:i], URLs[i+1:]...)
// 	return true
// }
func TestDeleteReturnsTrueIfDeleteIsPerformed(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	assert.True(t, deleted)
}

// same as previous
func TestDeleteReturnsFalseIfDeleteIsNotPerformed(t *testing.T) {
	URLs = []string{}
	deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	assert.False(t, deleted)
}

// import "errors"

// func Add(url string) (bool, error) {
// 	for _, u := range URLs {
// 		if u == url {
// 			return false, errors.New("url already exists.")
// 		}
// 	}

// 	URLs = append(URLs, url)
// 	return true, nil
// }
func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
	URLs = []string{"http://www.google.ca"}
	added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.False(t, added)
	assert.Equal(t, err.Error(), "url already exists.")
}

// Same as previous
func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
	URLs = []string{}
	added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.True(t, added)
	assert.Nil(t, err)
}
