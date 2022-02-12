package test

import (
	"github.com/Mausved/uniq/unique"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUniqWithoutParams(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamC(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamD(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`I love music.
I love music of Kartik.
I love music of Kartik.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamU(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`
Thanks.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamI(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamF(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`We love music.

I love music of Kartik.
Thanks.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}

func TestUniqWithParamS(t *testing.T) {
	opts := unique.Options{
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
	expected :=
		`I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`
	inputSlice := strings.Split(input, "\n")
	assert.Equal(t, expected, unique.Uniq(&inputSlice, opts))
}
