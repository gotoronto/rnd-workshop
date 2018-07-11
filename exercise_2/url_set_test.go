package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddANewURL(t *testing.T) {
	list := NewURLSet()
	list.Add("http://www.google.ca")

	assert.Equal(t, []string{"http://www.google.ca"}, list.URLs)
}

// func TestAddOnlyKeepsUniqueURL(t *testing.T) {
//	list := NewURLSet()
//	list.Add("http://www.google.ca")
//	list.Add("http://www.google.ca")

//	assert.Equal(t, []string{"http://www.google.ca"}, list.URLs)
// }

// func TestAddReturnsFalseAndErrorIfURLAlreadyExists(t *testing.T) {
//	list := NewURLSet("http://www.google.ca")
//	added, err := list.Add("http://www.google.ca")

//	assert.Equal(t, []string{"http://www.google.ca"}, list.URLs)
//	assert.False(t, added)
//	assert.Equal(t, err.Error(), "url already exists")
// }

// func TestAddReturnsTrueAndNilIfURLIsUnique(t *testing.T) {
//	list := NewURLSet()
//	added, err := list.Add("http://www.google.ca")

//	assert.Equal(t, []string{"http://www.google.ca"}, list.URLs)
//	assert.True(t, added)
//	assert.Nil(t, err)
// }
