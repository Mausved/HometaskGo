package unique

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWriteWithoutParams(t *testing.T) {
	const message = "Test without params."
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
	var output []string
	expected := append(output, inputRow)
	Append(&output, inputCurrentCount, inputRow, opts)
	require.Equal(t, expected, output, message)
}

func TestWriteWithUniqParamNotUniqRow(t *testing.T) {
	const message = "Test with uniq param. Not uniq input row."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        true,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 5
	inputRow := "not uniq row."
	var output []string
	var expected []string
	Append(&output, inputCurrentCount, inputRow, opts)
	require.Equal(t, expected, output, message)
}

func TestWriteWithUniqParamUniqRow(t *testing.T) {
	const message = "Test with uniq param. Uniq input row."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        true,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 1
	inputRow := "uniq row."
	var output []string
	expected := append(output, inputRow)
	Append(&output, inputCurrentCount, inputRow, opts)
	require.Equal(t, expected, output, message)
}

func TestWriteWithDoubleParamNotDoubleRow(t *testing.T) {
	const message = "Test with double param. Not double input row."
	opts := Options{
		Count:       false,
		Double:      true,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 1
	inputRow := "uniq row."
	var output []string
	var expected []string
	Append(&output, inputCurrentCount, inputRow, opts)
	require.Equal(t, expected, output, message)
}

func TestWriteWithDoubleParamDoubleRow(t *testing.T) {
	const message = "Test with double row. Double input row."
	opts := Options{
		Count:       false,
		Double:      true,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}
	inputCurrentCount := 5
	inputRow := "not uniq row."
	var output []string
	expected := append(output, inputRow)
	Append(&output, inputCurrentCount, inputRow, opts)
	require.Equal(t, expected, output, message)
}
