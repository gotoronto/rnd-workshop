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
	//Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

func TestAddOnlyKeepsUniqueURL(t *testing.T) {
	URLs = []string{}
	//Add("http://www.google.ca")
	//Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

// doc on deleting an element from an array: https://github.com/golang/go/wiki/SliceTricks#delete
func TestDeleteURL(t *testing.T) {
	URLs = []string{}
	//Add("http://www.google.ca")
	//Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
}

func TestDeleteDoesNothingIfURLNotExist(t *testing.T) {
	//Add("http://www.amazon.ca")
	//Delete("http://www.google.ca")

	assert.Equal(t, []string{"http://www.amazon.ca"}, URLs)
}

func TestDeleteReturnsTrueIfDeleteIsPerformed(t *testing.T) {
	URLs = []string{}
	//Add("http://www.google.ca")
	//deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	//assert.True(t, deleted)
}

func TestDeleteReturnsFalseIfDeleteIsNotPerformed(t *testing.T) {
	URLs = []string{}
	//deleted := Delete("http://www.google.ca")

	assert.Equal(t, []string{}, URLs)
	//assert.False(t, deleted)
}

func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
	URLs = []string{"http://www.google.ca"}
	//added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	//assert.False(t, added)
	//assert.Equal(t, err.Error(), "url already exists.")
}

func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
	URLs = []string{}
	//added, err := Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	//assert.True(t, added)
	//assert.Nil(t, err)
}
