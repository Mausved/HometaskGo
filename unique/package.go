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

func CountWords(currentString string, opts Options) string {
	words := strings.Split(currentString, Separator)
	minWords := Min(len(words), opts.NumFields)
	newString := strings.Join(words[minWords:], Separator)
	minChars := Min(len(newString), opts.NumChars)
	return newString[minChars:]
}

func Writer(currentCount int, output *string, row string, opts Options) {
	out := ""
	if opts.Count {
		out += strconv.Itoa(currentCount) + Separator
	}
	*output += out + row + "\n"
}
func writeRow(output *string, currentCount int, row string, opts Options) {
	if opts.Uniq {
		if currentCount == 1 {
			Writer(currentCount, output, row, opts)
		}
	} else if opts.Double && currentCount > 1 || !opts.Double {
		Writer(currentCount, output, row, opts)
	}
}

//func ToFormat(text string, opts Options) string {
//	formattedText := CountWords(text, opts)
//	if opts.Insensitive {
//		formattedText = strings.ToLower(formattedText)
//	}
//	return formattedText
//}

func Uniq(rows *[]string, opts Options) string {
	output := ""
	currentCount := 1
	currentRow := (*rows)[0]
	for idx, nextRow := range (*rows)[1:] {
		if strings.EqualFold(CountWords(nextRow, opts), CountWords(currentRow, opts)) {
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
