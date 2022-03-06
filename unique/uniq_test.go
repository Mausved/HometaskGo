package unique

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestUniqWithoutParams(t *testing.T) {
	const message = "Test without any params."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	input :=
		`I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`
	expected := []string{
		`I love music.`,
		``,
		`I love music of Kartik.`,
		`Thanks.`,
		`I love music of Kartik.`,
	}

	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamC(t *testing.T) {
	const message = "Test with print counts."
	opts := Options{
		Count:       true,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	input :=
		`I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`
	expected := []string{
		`3 I love music.`,
		`1 `,
		`2 I love music of Kartik.`,
		`1 Thanks.`,
		`2 I love music of Kartik.`,
	}

	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamD(t *testing.T) {
	const message = "Test with print doubles rows."
	opts := Options{
		Count:       false,
		Double:      true,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	input :=
		`I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`
	expected := []string{
		`I love music.`,
		`I love music of Kartik.`,
		`I love music of Kartik.`,
	}

	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamU(t *testing.T) {
	const message = "Test with print uniq rows."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        true,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	input :=
		`I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`
	expected := []string{
		``,
		`Thanks.`,
	}

	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamI(t *testing.T) {
	const message = "Test with ignore case parameter."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: true,
		NumFields:   0,
		NumChars:    0,
	}

	input :=
		`I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.`
	expected := []string{
		`I LOVE MUSIC.`,
		``,
		`I love MuSIC of Kartik.`,
		`Thanks.`,
		`I love music of kartik.`,
	}

	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamF(t *testing.T) {
	const message = "Test with param to ignore 1 word."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   1,
		NumChars:    0,
	}

	input :=
		`We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`
	expected := []string{
		`We love music.`,
		``,
		`I love music of Kartik.`,
		`Thanks.`,
	}
	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqWithParamS(t *testing.T) {
	const message = "Test with param to ignore 1 char."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    1,
	}

	input :=
		`I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`
	expected := []string{
		`I love music.`,
		``,
		`I love music of Kartik.`,
		`We love music of Kartik.`,
		`Thanks.`,
	}
	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqOneRowInput(t *testing.T) {
	const message = "Test with one row."
	opts := Options{
		Count:       false,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	input := "1"
	expected := []string{"1"}
	inputSlice := strings.Split(input, "\n")
	actual, err := Uniq(inputSlice, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}

func TestUniqEmptyInput(t *testing.T) {
	const message = "Test with empty input."
	opts := Options{
		Count:       true,
		Double:      false,
		Uniq:        false,
		Insensitive: false,
		NumFields:   0,
		NumChars:    0,
	}

	var input []string
	var expected []string = nil
	actual, err := Uniq(input, opts)
	require.Equal(t, expected, actual, message)
	require.Equal(t, nil, err)
}
