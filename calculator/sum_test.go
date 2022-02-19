package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSums(t *testing.T) {
	require.Equal(t, float64(-2), Calculator("-2"), "-2=-2")
	require.Equal(t, float64(0), Calculator("(-2)+2"), "-2+2=0")
	require.Equal(t, float64(0), Calculator("2+(-2)"), "2-2=0")
	require.Equal(t, float64(0), Calculator("2+-2"), "2-2=0")
	require.Equal(t, float64(4), Calculator("2--2"), "2+2=4")
	require.Equal(t, float64(0), Calculator("2---2"), "2-2=0")
}
