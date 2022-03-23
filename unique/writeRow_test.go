package unique

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestWriteRowWithoutCount(t *testing.T) {
	const message = "Test without print counts."
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
	require.Equal(t, expected, writeRow(inputCurrentCount, inputRow, opts), message)
}

func TestWriteRowWithCount(t *testing.T) {
	const message = "Test with print counts."
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
	require.Equal(t, expected, writeRow(inputCurrentCount, inputRow, opts), message)
}
