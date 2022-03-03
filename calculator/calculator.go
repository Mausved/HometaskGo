package calculator

import (
	"github.com/Mausved/calculator/stack"
	"github.com/cstockton/go-conv"
	"math"
	"regexp"
	"strings"
)

func Calculator(input string) float64 {
	isRestricted, err := regexp.MatchString(restrictedSymbolsPattern, input)
	success, formatted := toFormat(input)
	if !success || err != nil || isRestricted || formatted == "" {
		return math.NaN()
	}
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
	return math.Round(CountReversePolishNotation(polish)*100) / 100
}

const (
	minus         = "-"
	plus          = "+"
	multiply      = "*"
	div           = "/"
	operators     = minus + plus + multiply + div
	openBrackets  = "({["
	closeBrackets = ")}]"
	brackets      = openBrackets + closeBrackets
	digits        = "0123456789"
	notDigits     = operators + brackets
	dot           = "."
)

var escapedOperators = "\\" + strings.Join(strings.Split(operators, ""), "\\")
var escapedBrackets = "\\" + strings.Join(strings.Split(brackets, ""), "\\")
var allowedSymbols = digits + escapedOperators + escapedBrackets + dot
var restrictedSymbolsPattern = "[^" + allowedSymbols + "]"

func Sum(operands ...interface{}) float64 {
	result := 0.
	for _, number := range operands {
		val, err := conv.Float64(number)
		if err != nil {
			return math.NaN()
		}
		result += val
	}
	return result
}
