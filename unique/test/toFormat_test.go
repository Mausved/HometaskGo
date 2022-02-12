package test

import (
	"github.com/Mausved/uniq/unique"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestWithoutFormatting(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	input := "uniq first row."
	expected := input
	assert.Equal(t, expected, unique.ToFormat(input, opts))
}

func TestWithCountingFormatting(t *testing.T) {
	const ignoredWords = 2
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   ignoredWords,
		NumChars:    0,
	}
	input := "in this test will ignored the first two begin words."
	expected := strings.Join(strings.Split(input, " ")[ignoredWords:], " ")
	assert.Equal(t, expected, unique.ToFormat(input, opts))
}

func TestWithCountingFormattingIgnoredAllWords(t *testing.T) {
	const ignoredWords = 200
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   ignoredWords,
		NumChars:    0,
	}
	input := "so little words in compared with constant at 39 row..."
	expected := ""
	assert.Equal(t, expected, unique.ToFormat(input, opts))
}

func TestWithRegisterFormatting(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: true,
		NumFields:   0,
		NumChars:    0,
	}
	input := "THIS TEXT WILL BE IN LOWER"
	expected := strings.ToLower(input)
	assert.Equal(t, expected, unique.ToFormat(input, opts))
}

func TestRegisterAndCountingFormatting(t *testing.T) {
	const ignoredWords = 5
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: true,
		NumFields:   ignoredWords,
		NumChars:    0,
	}
	input := "IN THIS TEST WILL BE IN LOW AND IGNORED THE FIRST FIVE WORDS"
	expected := strings.Join(strings.Split(strings.ToLower(input), " ")[ignoredWords:], " ")
	assert.Equal(t, expected, unique.ToFormat(input, opts))
}
