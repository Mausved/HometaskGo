package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestStandardMultiply(t *testing.T) {
	require.Equal(t, 5., Calculator("1.25*4"), "1.25*4=5")
	require.Equal(t, 3.33, Calculator("1.111*3"), "1.111*3=3.33")
	require.Equal(t, 0., Calculator("-2*0"), "-2*0=0")
	require.Equal(t, 0., Calculator("2*0"), "2*0=0")
	require.Equal(t, 4., Calculator("2*2"), "2*2=4")
	require.Equal(t, -4., Calculator("2*(-2)"), "2*(-2)=4")
	require.Equal(t, 12., Calculator("2*2*3"), "2*2*3=12")
	require.Equal(t, 120., Calculator("1*2*3*4*5"), "1*2*3*4*5=120")
}

func TestNaNMultiply(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2*---2")), "2*---2=2*-2=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("0.1****3")), "0.1****3=0.1*3=0.3")
	require.Equal(t, true, math.IsNaN(Calculator("2****3")), "2****3=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2**0**3")), "2**0**3=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2*0**3")), "2*0**3=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2****3")), "2****3=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2*-2")), "2*-2=NaN")
}
