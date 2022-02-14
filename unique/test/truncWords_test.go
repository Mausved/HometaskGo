package test

import (
	"github.com/Mausved/uniq/unique"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithoutFields(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	input := "there're three words."
	expected := input
	assert.Equal(t, expected, unique.TruncWords(input, opts))
}

func TestTrunc1Word(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   1,
		NumChars:    0,
	}
	input := "there're three words."
	expected := "three words."

	assert.Equal(t, expected, unique.TruncWords(input, opts))
}

func TestTruncAll(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   3,
		NumChars:    0,
	}
	input := "there're three words."
	expected := ""

	assert.Equal(t, expected, unique.TruncWords(input, opts))
}

func TestNumFieldsMoreThanNumberOfWords(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   300,
		NumChars:    0,
	}
	input := "this test is check correctly behavior NumFields more the number of words in this string."
	expected := ""

	assert.Equal(t, expected, unique.TruncWords(input, opts))
}
