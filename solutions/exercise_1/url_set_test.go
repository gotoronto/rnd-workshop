package urlset

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
	var added bool
	var err error
	added, err = Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.False(t, added)
	assert.EqualError(t, err, "url already exists.")
}

func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
	URLs = []string{}
	var added bool
	var err error
	added, err = Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, URLs)
	assert.True(t, added)
	assert.Nil(t, err)
}
