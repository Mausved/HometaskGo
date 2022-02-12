package unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithoutFields(t *testing.T) {
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	input := "there're three words."
	expected := input
	assert.Equal(t, expected, countWords(input, opts))
}

func TestEscape1Word(t *testing.T) {
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   1,
		NumChars:    0,
	}
	input := "there're three words."
	expected := "three words."

	assert.Equal(t, expected, countWords(input, opts))
}

func TestEscapeAll(t *testing.T) {
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   3,
		NumChars:    0,
	}
	input := "there're three words."
	expected := ""

	assert.Equal(t, expected, countWords(input, opts))
}

func TestNumFieldsMoreThanNumberOfWords(t *testing.T) {
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   300,
		NumChars:    0,
	}
	input := "this test is check correctly behavior NumFields more the number of words in this string."
	expected := ""

	assert.Equal(t, expected, countWords(input, opts))
}
