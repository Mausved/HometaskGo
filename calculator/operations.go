package calculator

import (
	"math"
	"strings"
)

func getOperation(x float64, y float64, operator string) float64 {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		if y == 0 {
			return math.NaN()
		}
		return x / y
	default:
		return math.NaN()
	}
}

func isOperator(symbol string) bool {
	return strings.ContainsAny(symbol, operators) && len(symbol) == 1
}

func isOpenBracket(symbol string) bool {
	return strings.ContainsAny(symbol, openBrackets) && len(symbol) == 1
}

func isCloseBracket(symbol string) bool {
	return strings.ContainsAny(symbol, closeBrackets) && len(symbol) == 1
}

func isDigit(symbol string) bool {
	return strings.ContainsAny(symbol, digits) && len(symbol) == 1
}

const (
	maxPriority = 3
	midPriority = 2
	minPriority = 1
)

func getPriority(operator string) int {
	if operator == multiply || operator == div {
		return maxPriority
	}
	if operator == minus || operator == plus {
		return midPriority
	}
	return minPriority
}
