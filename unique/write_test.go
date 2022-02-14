package unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteWithoutParams(t *testing.T) {
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
	Write(&output, inputCurrentCount, inputRow, opts)
	assert.Equal(t, expected, output)
}

func TestWriteWithUniqParamNotUniqRow(t *testing.T) {
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
	Write(&output, inputCurrentCount, inputRow, opts)
	assert.Equal(t, expected, output)
}

func TestWriteWithUniqParamUniqRow(t *testing.T) {
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
	Write(&output, inputCurrentCount, inputRow, opts)
	assert.Equal(t, expected, output)
}

func TestWriteWithDoubleParamNotDoubleRow(t *testing.T) {
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
	Write(&output, inputCurrentCount, inputRow, opts)
	assert.Equal(t, expected, output)
}

func TestWriteWithDoubleParamDoubleRow(t *testing.T) {
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
	Write(&output, inputCurrentCount, inputRow, opts)
	assert.Equal(t, expected, output)
}
