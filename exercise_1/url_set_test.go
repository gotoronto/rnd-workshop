package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Uncomment one test at a time and write code in crawler/crawler.go to satisfy the test you just uncommented.

// To run the test
// dep ensure     // to install dependencies
// go test ./...

// Resources:
// https://gobyexample.com/
// https://golang.org/doc/ (use the search bar for doc)
// Google and Stack Overflow are your friends :)

// =============================================
// TEST CASES START HERE
// =============================================

func TestAddANewURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// func TestAddOnlyKeepsUniqueURL(t *testing.T) {
//	URLs = []string{}
//	Add("http://www.google.ca")
//	Add("http://www.google.ca")

//	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
// }

// func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
//	URLs = []string{"http://www.google.ca"}
//	added, err := Add("http://www.google.ca")

assert.Error(t, err)
//	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
//	assert.False(t, added)
//	assert.Equal(t, err.Error(), "url already exists.")
// }

// func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
//	URLs = []string{}
//	added, err := Add("http://www.google.ca")

//	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
//	assert.True(t, added)
//	assert.Nil(t, err)
// }
