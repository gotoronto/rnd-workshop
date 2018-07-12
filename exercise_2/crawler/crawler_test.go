package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddANewURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

func TestAddOnlyKeepsUniqueURL(t *testing.T) {
	URLs = []string{}
	Add("http://www.google.ca")
	Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
}

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
