package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestOkBrackets(t *testing.T) {
	require.Equal(t, 0., Calculator("()"), "() = 0")
	require.Equal(t, 4., Calculator("(2+2)"), "(2+2) = 4")
	require.Equal(t, 3., Calculator("(1)+(2)"), "(1)+(2) = 3")
	require.Equal(t, 0., Calculator("(2)+(-2)"), "(2)+(-2) = 0")
	require.Equal(t, 0., Calculator("(2)-(2)"), "(2)-(2) = 0")
	require.Equal(t, 0., Calculator("(2)-2"), "(2)-2 = 0")
	require.Equal(t, 0., Calculator("-(2)+2"), "-(2)+2 = 0")
	require.Equal(t, 0., Calculator("(-2)+2"), "(-2)+2 = 0")

}

func TestInvalidBrackets(t *testing.T) {
	const message = "invalid brackets returned NaN"
	require.Equal(t, true, math.IsNaN(Calculator(")(")), message)
	require.Equal(t, true, math.IsNaN(Calculator("()()()()())()()()()(")), message)
	require.Equal(t, true, math.IsNaN(Calculator(")))(((()()()(")), message)
	require.Equal(t, true, math.IsNaN(Calculator("2-2()")), "2-2() = NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2-2(2)")), "2-2(2) = NaN")
}
