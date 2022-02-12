package unique

import (
	"strconv"
	"strings"
)

const separator = " "

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

func countWords(currentString string, opts Options) string {
	words := strings.Split(currentString, separator)
	minWords := min(len(words), opts.NumFields)
	newString := strings.Join(words[minWords:], separator)
	minChars := min(len(newString), opts.NumChars)
	return newString[minChars:]
}

func writer(currentCount int, output *string, row string, opts Options) {
	out := ""
	if opts.Count {
		out += strconv.Itoa(currentCount) + separator
	}
	*output += out + row + "\n"
}
func writeRow(output *string, currentCount int, row string, opts Options) {
	if opts.Uniq {
		if currentCount == 1 {
			writer(currentCount, output, row, opts)
		}
	} else if opts.Double && currentCount > 1 || !opts.Double {
		writer(currentCount, output, row, opts)
	}
}

func toFormat(text string, opts Options) string {
	formattedText := countWords(text, opts)
	if opts.Insensitive {
		formattedText = strings.ToLower(formattedText)
	}
	return formattedText
}

func Uniq(rows *[]string, opts Options) string {
	output := ""
	currentCount := 1
	currentRow := (*rows)[0]
	for idx, nextRow := range (*rows)[1:] {
		formattedCurrentRow := toFormat(currentRow, opts)
		formattedNextRow := toFormat(nextRow, opts)
		if formattedNextRow == formattedCurrentRow {
			currentCount++
		} else {
			writeRow(&output, currentCount, currentRow, opts)
			currentRow = nextRow
			currentCount = 1
		}
		if idx+1 == len(*rows)-1 {
			writeRow(&output, currentCount, currentRow, opts)
		}
	}
	return output[:len(output)-1]
}
