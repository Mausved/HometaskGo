package test

import (
	"github.com/Mausved/uniq/unique"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWriteRowWithoutCount(t *testing.T) {
	opts := unique.Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 1
	textResult := ""
	inputRow := "uniq first row."
	expected := inputRow + "\n"
	unique.Writer(inputCurrentCount, &textResult, inputRow, opts)
	assert.Equal(t, expected, textResult)
}

func TestWriteRowWithCount(t *testing.T) {
	opts := unique.Options{
		Count:       true,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 5
	textResult := ""
	inputRow := "row with five repeats."
	expected := strconv.Itoa(inputCurrentCount) + unique.Separator + inputRow + "\n"
	unique.Writer(inputCurrentCount, &textResult, inputRow, opts)
	assert.Equal(t, expected, textResult)
}
