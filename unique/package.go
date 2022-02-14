package unique

import (
	"strconv"
	"strings"
)

const Separator = " "

type Options struct {
	Count       bool
	Double      bool
	Uniq        bool
	Insensitive bool
	NumFields   int
	NumChars    int
}

func Min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func TruncWords(currentString string, opts Options) string {
	words := strings.Split(currentString, Separator)
	minWords := Min(len(words), opts.NumFields)
	newString := strings.Join(words[minWords:], Separator)
	minChars := Min(len(newString), opts.NumChars)
	return newString[minChars:]
}

func WriteRow(currentCount int, row string, opts Options) string {
	out := ""
	if opts.Count {
		out = strconv.Itoa(currentCount) + Separator
	}
	return out + row
}

func Write(output *[]string, currentCount int, row string, opts Options) {
	if opts.Uniq {
		if currentCount == 1 {
			*output = append(*output, WriteRow(currentCount, row, opts))

		}
	} else if opts.Double && currentCount > 1 || !opts.Double {
		*output = append(*output, WriteRow(currentCount, row, opts))
	}
}

func Uniq(rows *[]string, opts Options) string {
	if len(*rows) == 0 {
		return ""
	}

	output := make([]string, 0, len(*rows))
	currentCount := 1
	currentRow := (*rows)[0]

	if len(*rows) == 1 {
		Write(&output, currentCount, currentRow, opts)
		return strings.Join(output, "\n")
	}

	for idx, nextRow := range (*rows)[1:] {
		if strings.EqualFold(TruncWords(nextRow, opts), TruncWords(currentRow, opts)) {
			currentCount++
		} else {
			Write(&output, currentCount, currentRow, opts)
			currentRow = nextRow
			currentCount = 1
		}
		if idx+1 == len(*rows)-1 {
			Write(&output, currentCount, currentRow, opts)
		}
	}
	return strings.Join(output, "\n")
}
