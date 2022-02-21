package calculator

import (
	"github.com/Mausved/calculator/stack"
	"strconv"
	"strings"
)

func CountReversePolishNotation(polish *[]string) float64 {
	floatStack := stack.Stack{}
	for _, element := range *polish {
		if !isOperator(element) {
			number, _ := strconv.ParseFloat(element, 64)
			floatStack.Push(number)
		} else {
			RightOperand := floatStack.GetTop().(float64)
			floatStack.Pop()
			if floatStack.IsEmpty() {
				floatStack.Push(float64(0))
			}
			LeftOperand := floatStack.GetTop().(float64)
			floatStack.Pop()
			operator := element
			floatStack.Push(getOperation(LeftOperand, RightOperand, operator))
		}
	}
	return Sum(floatStack.PourStackToSlice()...)
}

func makeNumber(str string) string {
	nextNotDigitPos := strings.IndexAny(str, notDigits)
	if nextNotDigitPos == -1 {
		nextNotDigitPos = len(str)
	}
	number := str[0:nextNotDigitPos]
	return number
}

func makeOperator(operator string, polish *[]string, stringStack *stack.Stack) {
	for !stringStack.IsEmpty() && getPriority(stringStack.GetTop().(string)) >= getPriority(operator) {
		top := stringStack.GetTop().(string)
		*polish = append(*polish, top)
		stringStack.Pop()
	}
	stringStack.Push(operator)
}

func makeExpressionInBrackets(stringStack *stack.Stack, polish *[]string) {
	top := stringStack.GetTop().(string)
	for !isOpenBracket(top) {
		*polish = append(*polish, top)
		stringStack.Pop()
		top = stringStack.GetTop().(string)
	}
	stringStack.Pop()
}
