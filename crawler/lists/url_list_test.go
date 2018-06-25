package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddANewURL(t *testing.T) {
	list := NewURLList()
	list.Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, list.urls)
}

func TestAddOnlyKeepsUniqueURL(t *testing.T) {
	list := NewURLList()
	list.Add("http://www.google.ca")
	list.Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, list.urls)
}

func TestDeleteURL(t *testing.T) {
	list := NewURLList()
	list.Add("http://www.google.ca")
	list.Delete("http://www.google.ca")

	assert.Equal(t, []string{}, list.urls)
}

func TestDeleteDoesNothingIfURLNotExist(t *testing.T) {
	list := NewURLList()
	list.Add("http://www.amazon.ca")
	list.Delete("http://www.google.ca")

	assert.Equal(t, []string{"http://www.amazon.ca"}, list.urls)
}

func TestDeleteReturnsTrueIfDeleteIsPerformed(t *testing.T) {
	list := NewURLList()
	list.Add("http://www.google.ca")
	deleted := list.Delete("http://www.google.ca")

	assert.Equal(t, []string{}, list.urls)
	assert.True(t, deleted)
}

func TestDeleteReturnsFalseIfDeleteIsNotPerformed(t *testing.T) {
	list := NewURLList()
	deleted := list.Delete("http://www.google.ca")

	assert.Equal(t, []string{}, list.urls)
	assert.False(t, deleted)
}

func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
	list := NewURLList("http://www.google.ca")
	added, err := list.Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, list.urls)
	assert.False(t, added)
	assert.Equal(t, err.Error(), "url already exists")
}

func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
	list := NewURLList()
	added, err := list.Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, list.urls)
	assert.True(t, added)
	assert.Nil(t, err)
}
