package calculator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitMultiply(t *testing.T) {
	require.Equal(t, float64(0), Calculator("-2*0"), "-2*0=0")
	require.Equal(t, float64(0), Calculator("2*0"), "2*0=0")
	require.Equal(t, float64(4), Calculator("2*2"), "2*2=4")
	require.Equal(t, float64(-4), Calculator("2*(-2)"), "2*(-2)=4")
	require.Equal(t, float64(-4), Calculator("2*-2"), "2*(-2)=4")
	require.Equal(t, float64(12), Calculator("2*2*3"), "2*2*3=12")
	require.Equal(t, float64(120), Calculator("1*2*3*4*5"), "1*2*3*4*5=120")
}
