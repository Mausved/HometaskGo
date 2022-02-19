package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSums(t *testing.T) {
	require.Equal(t, float64(-2), Calculator("-2"), "-2=-2")
	require.Equal(t, float64(0), Calculator("(-2)+2"), "-2+2=0")
	require.Equal(t, float64(0), Calculator("2+(-2)"), "2-2=0")

}

func TestNanSums(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2+-2")), "2/0=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2--2")), "2--2=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2---2")), "2---2=NaN")
	//
	//require.Equal(t, float64(0), Calculator("2+-2"), "2-2=0")
	//require.Equal(t, float64(4), Calculator("2--2"), "2+2=4")
	//require.Equal(t, float64(0), Calculator("2---2"), "2-2=0")
}
