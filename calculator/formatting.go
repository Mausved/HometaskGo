package calculator

import (
	"strings"
)

func checkLast(lastSymbol string, operator string) string {
	checked := ""
	if isOperator(lastSymbol) {
		checked = lastSymbol
	} else if isOpenBracket(lastSymbol) {
		checked = operator
	}
	if isOperator(lastSymbol) || isOpenBracket(lastSymbol) {
		switch checked {
		case minus, plus:
			return "0"
		case multiply, div:
			return "1"
		}
	}
	return ""
}

func toFormat(str *string) string {
	brackets := 0
	prevOperator := ""
	format := ""
	slice := strings.Split(*str, "")
	for idx, word := range slice {
		if isOperator(word) {
			currOperator := word
			if idx-1 > 0 && isOperator(slice[idx-1]) {
				return ""
			}
			if prevOperator != word {
				format += prevOperator
				prevOperator = currOperator
			}
		} else {
			if len(format) > 0 {
				lastSymbol := format[len(format)-1]
				format += checkLast(string(lastSymbol), prevOperator)
			}
			format += prevOperator + word
			prevOperator = ""
			if isOpenBracket(word) {
				if idx-1 > 0 && isDigit(slice[idx-1]) {
					return ""
				}
				brackets++
			} else if isCloseBracket(word) {
				brackets--
			}
			if brackets < 0 {
				return ""
			}
		}
	}
	if brackets != 0 {
		return ""
	}
	return format
}
