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

func truncWords(currentString string, opts Options) string {
	words := strings.Split(currentString, Separator)
	minWords := Min(len(words), opts.NumFields)
	newString := strings.Join(words[minWords:], Separator)
	minChars := Min(len(newString), opts.NumChars)
	return newString[minChars:]
}

func writeRow(currentCount int, currentRow string, opts Options) string {
	out := ""
	if opts.Count {
		out = strconv.Itoa(currentCount) + Separator
	}
	return out + currentRow
}

func addRow(result []string, currentCount int, currentRow string, opts Options) []string {
	if opts.Uniq {
		if currentCount == 1 {
			return append(result, writeRow(currentCount, currentRow, opts))
		}
	} else if !opts.Double || opts.Double && currentCount > 1 {
		return append(result, writeRow(currentCount, currentRow, opts))
	}
	return result
}

func Uniq(rows []string, opts Options) string {
	if len(rows) == 0 {
		return ""
	}

	result := make([]string, 0, len(rows))
	currentCount := 1
	currentRow := rows[0]

	if len(rows) == 1 {
		result = addRow(result, currentCount, currentRow, opts)
		return result[0]
	}

	for idx, nextRow := range rows[1:] {
		if strings.EqualFold(truncWords(nextRow, opts), truncWords(currentRow, opts)) {
			currentCount++
		} else {
			result = addRow(result, currentCount, currentRow, opts)
			currentRow = nextRow
			currentCount = 1
		}
		if idx+1 == len(rows)-1 {
			result = addRow(result, currentCount, currentRow, opts)
		}
	}
	return strings.Join(result, "\n")
}
