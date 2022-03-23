package unique

import (
	"errors"
	"strconv"
	"strings"
)

const (
	Separator  = " "
	ErrorFlags = "multiply flag using"
)

type Options struct {
	Count       bool
	Double      bool
	Uniq        bool
	Insensitive bool
	NumFields   int
	NumChars    int
}

func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

func truncRow(currentString string, opts Options) string {
	words := strings.Fields(currentString)
	minWords := min(len(words), opts.NumFields)
	newString := strings.Join(words[minWords:], Separator)
	minChars := min(len(newString), opts.NumChars)
	return newString[minChars:]
}

func writeRow(currentCount int, currentRow string, opts Options) string {
	out := ""
	if opts.Count {
		out = strconv.Itoa(currentCount) + Separator
	}
	return out + currentRow
}

func addRow(result []string, currentCount int, currentRow string, simpleMode bool, opts Options) []string {
	if simpleMode || opts.Uniq && currentCount == 1 || opts.Double && currentCount > 1 {
		return append(result, writeRow(currentCount, currentRow, opts))
	}
	return result
}

func Uniq(rows []string, opts Options) ([]string, error) {
	if len(rows) == 0 {
		return nil, nil
	}

	if opts.Uniq && opts.Double || opts.Uniq && opts.Count || opts.Double && opts.Count {
		return nil, errors.New(ErrorFlags)
	}

	result := make([]string, 0)
	currentCount := 1
	currentRow := rows[0]
	simpleMode := !opts.Double && !opts.Uniq

	if len(rows) == 1 {
		result = addRow(result, currentCount, currentRow, simpleMode, opts)
		return result, nil
	}

	for idx, nextRow := range rows[1:] {
		if strings.EqualFold(truncRow(nextRow, opts), truncRow(currentRow, opts)) {
			currentCount++
		} else {
			result = addRow(result, currentCount, currentRow, simpleMode, opts)
			currentRow = nextRow
			currentCount = 1
		}
		if idx+1 == len(rows)-1 {
			result = addRow(result, currentCount, currentRow, simpleMode, opts)
		}
	}
	return result, nil
}
