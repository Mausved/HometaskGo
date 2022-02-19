package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitTestSumAndMultiply(t *testing.T) {
	require.Equal(t, float64(9), Calculator("(1+2)*3"), "(1+2)*3 = 9")
	require.Equal(t, float64(-5), Calculator("1+2*(-3)"), "1+2*(-3) = 7")
	require.Equal(t, float64(-1), Calculator("(1+2)-4"), "(1+2)-4 = -1")
}
