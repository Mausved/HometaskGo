package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitTestOnlySum(t *testing.T) {
	const message = "(1+2)-3 = 0"
	input := "(1+2)-3"
	expected := "0"
	require.Equal(t, expected, Calculator(input), message)
}

func TestInitTestSumAndMultiply(t *testing.T) {
	const message = "(1+2)*3 = 9"
	input := "(1+2)*3"
	expected := "9"
	require.Equal(t, expected, Calculator(input), message)
}
