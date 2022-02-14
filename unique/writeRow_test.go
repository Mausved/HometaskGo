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
	inputRow := "uniq first row."
	expected := inputRow
	assert.Equal(t, expected, WriteRow(inputCurrentCount, inputRow, opts))
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
	inputRow := "row with five repeats."
	expected := strconv.Itoa(inputCurrentCount) + Separator + inputRow
	assert.Equal(t, expected, WriteRow(inputCurrentCount, inputRow, opts))
}
