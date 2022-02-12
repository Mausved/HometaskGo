package unique

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestWriteRowWithoutCount(t *testing.T) {
	opts := Options{
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
	writer(inputCurrentCount, &textResult, inputRow, opts)
	assert.Equal(t, expected, textResult)
}

func TestWriteRowWithCount(t *testing.T) {
	opts := Options{
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
	expected := strconv.Itoa(inputCurrentCount) + separator + inputRow + "\n"
	writer(inputCurrentCount, &textResult, inputRow, opts)
	assert.Equal(t, expected, textResult)
}
