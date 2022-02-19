package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestInitDiv(t *testing.T) {
	require.Equal(t, 1., Calculator("2/2"), "2/2=1")
	require.Equal(t, 1.5, Calculator("3/2"), "3/2=1.5")
	require.Equal(t, 2., Calculator("4/2"), "4/2=2")
	require.Equal(t, 10., Calculator("100/10"), "100/10=10")
	require.Equal(t, 0.5, Calculator("24.05/47.69"), "24.05/47.69=0.5")
	require.Equal(t, 0.51, Calculator("24.05/47.69+0.001"), "24.05/47.69+0.001=0.51")
	require.Equal(t, -1., Calculator("-2/2"), "-2/2=-1")

}

func TestNaNDiv(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2/0")), "2/0=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("-2/0")), "-2/0=NaN")
	require.Equal(t, false, math.IsNaN(Calculator("2/-2")), "2/-2=NaN")
}
