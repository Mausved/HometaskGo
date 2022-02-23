package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWithoutFields(t *testing.T) {
	const message = "Do not trunc text."
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
	require.Equal(t, expected, truncWords(input, opts), message)
}

func TestTrunc1Word(t *testing.T) {
	const message = "Test with only one word trunc."
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

	require.Equal(t, expected, truncWords(input, opts), message)
}

func TestTruncAll(t *testing.T) {
	const message = "Test with num fields equal to number of words in text."
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

	require.Equal(t, expected, truncWords(input, opts), message)
}

func TestNumFieldsMoreThanNumberOfWords(t *testing.T) {
	const message = "Test with num fields more than number of words in input text."
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

	require.Equal(t, expected, truncWords(input, opts), message)
}
