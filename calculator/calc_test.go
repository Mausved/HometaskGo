package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitTestSumAndMultiply(t *testing.T) {
	require.Equal(t, 9., Calculator("(1+2)*3"), "(1+2)*3 = 9")
	require.Equal(t, -5., Calculator("1+2*(-3)"), "1+2*(-3) = 7")
	require.Equal(t, -1., Calculator("(1+2)-4"), "(1+2)-4 = -1")
	require.Equal(t, 16., Calculator("(2+2)*4"), "(2+2)*4 = 16")
	require.Equal(t, 10., Calculator("2+2*4"), "2+2*4 = 10")
	require.Equal(t, 0., Calculator("(2-2)*4"), "(2-2)*4 = 0")
}
