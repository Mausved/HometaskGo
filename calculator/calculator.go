package calculator

import (
	"github.com/Mausved/uniq/stack"
	"math"
	"regexp"
)

func Calculator(input string) float64 {
	isRestricted, err := regexp.MatchString(restrictedSymbolsPattern, input)
	formatted := toFormat(&input)
	if err != nil || isRestricted || formatted == "" {
		return math.NaN()
	}
	//fmt.Println(fo)
	stringStack := stack.Stack{}
	var polish []string
	for i := 0; i < len(formatted); {
		symbol := string(formatted[i])
		if isDigit(symbol) {
			number := makeNumber(formatted[i:])
			polish = append(polish, number)
			i += len(number)
		} else {
			if isOperator(symbol) {
				makeOperator(symbol, &polish, &stringStack)
			} else if isOpenBracket(symbol) {
				stringStack.Push(symbol)
			} else if isCloseBracket(symbol) {
				makeExpressionInBrackets(&stringStack, &polish)
			}
			i++
		}
	}
	for !stringStack.IsEmpty() {
		polish = append(polish, stringStack.GetTop().(string))
		stringStack.Pop()
	}
	return CountReversePolishNotation(&polish)
}

const (
	restrictedSymbols        = "a-zA-Z,!@#$%^&~=/|\\;:<>?"
	restrictedSymbolsPattern = "[" + restrictedSymbols + "]"
	minus                    = "-"
	plus                     = "+"
	multiply                 = "*"
	div                      = "/"
	operators                = minus + plus + multiply + div
	openBrackets             = "({["
	closeBrackets            = ")}]"
	brackets                 = openBrackets + closeBrackets
	digits                   = "0123456789"
	notDigits                = operators + brackets
)

func Sum(operands ...interface{}) float64 {
	result := float64(0)
	for _, number := range operands {
		val, ok := number.(float64)
		if !ok {
			val = 0
		}
		result += val
	}
	return result
}
