package calculator

import (
	"strings"
)

func checkMinus(operator *string, quantity int) {
	if *operator == minus && quantity%2 == 0 {
		*operator = plus
	}
}

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
	countCurrOperator := 0
	format := ""
	for _, symbol := range strings.Split(*str, "") {
		if isOperator(symbol) {
			operator := symbol
			if currOperator == symbol {
				countCurrOperator++
			} else {
				checkMinus(&currOperator, countCurrOperator)
				format += currOperator
				currOperator = operator
				countCurrOperator = 1
			}
		} else {
			checkMinus(&currOperator, countCurrOperator)
			if len(format) > 0 {
				lastSymbol := format[len(format)-1]
				format += checkLast(string(lastSymbol), currOperator)
			}
			format += currOperator + symbol

			currOperator = ""
			countCurrOperator = 0

			if isOpenBracket(symbol) {
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
	return format
}
