package calculator

import (
	"strings"
)

func checkLast(lastSymbol string, operator string) string {
	if operator == minus && !isDigit(lastSymbol) {
		return "0"
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
				return false, ""
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
					return false, ""
				}
				brackets++
			} else if isCloseBracket(word) {
				brackets--
			}
			if brackets < 0 {
				return false, ""
			}
		}
	}
	if brackets != 0 {
		return false, ""
	}
	return true, format
}
