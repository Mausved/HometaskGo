package calculator

import (
	"fmt"
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
	currOperator := ""
	format := ""
	slice := strings.Split(*str, "")
	fmt.Println(*str)
	for idx, symbol := range slice {
		if isOperator(symbol) {
			operator := symbol
			if idx-1 > 0 && isOperator(slice[idx-1]) {
				return ""
			}
			if currOperator != symbol {
				format += currOperator
				currOperator = operator
			}
		} else {
			if len(format) > 0 {
				lastSymbol := format[len(format)-1]
				format += checkLast(string(lastSymbol), currOperator)
			}
			format += currOperator + symbol

			currOperator = ""

			if isOpenBracket(symbol) {
				if idx-1 > 0 && isDigit(slice[idx-1]) {
					return ""
				}
				brackets++
			} else if isCloseBracket(symbol) {
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
	fmt.Println("formatted=", format)
	return format
}
