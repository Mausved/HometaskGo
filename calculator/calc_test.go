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
	require.Equal(t, true, math.IsNaN(Calculator("((((((((((")), message)
	require.Equal(t, true, math.IsNaN(Calculator(")(")), message)
	require.Equal(t, true, math.IsNaN(Calculator("()()()()())()()()()(")), message)
	require.Equal(t, true, math.IsNaN(Calculator(")))(((()()()(")), message)
	require.Equal(t, true, math.IsNaN(Calculator("2-2()")), "2-2() = NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2-2(2)")), "2-2(2) = NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2*-2")), "2*-2 = NaN")
}

func TestSums(t *testing.T) {
	require.Equal(t, float64(-2), Calculator("-2"), "-2=-2")
	require.Equal(t, float64(0), Calculator("(-2)+2"), "(-2)+2=0")
	require.Equal(t, float64(0), Calculator("2+(-2)"), "2+(-2)=0")
}

func TestNanSums(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2+-2")), "2+-0=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2--2")), "2--2=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2---2")), "2---2=NaN")
}

func TestStandardMultiply(t *testing.T) {
	require.Equal(t, 5., Calculator("1.25*4"), "1.25*4=5")
	require.Equal(t, 3.333, Calculator("1.111*3"), "1.111*3=3.33")
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

func TestDiv(t *testing.T) {
	require.Equal(t, 1., Calculator("2/2"), "2/2=1")
	require.Equal(t, 1.5, Calculator("3/2"), "3/2=1.5")
	require.Equal(t, 2., Calculator("4/2"), "4/2=2")
	require.Equal(t, 10., Calculator("100/10"), "100/10=10")
	require.Equal(t, 0.504298595093311, Calculator("24.05/47.69"), "24.05/47.69=0.5")
	require.Equal(t, 0.505298595093311, Calculator("24.05/47.69+0.001"), "24.05/47.69+0.001=0.51")
	require.Equal(t, -1., Calculator("-2/2"), "-2/2=-1")
	require.Equal(t, -1., Calculator("2/(-2)"), "2/(-2)=1")
}

func TestNaNDiv(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2/0")), "2/0=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("-2/0")), "-2/0=NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2/-2")), "2/-2=NaN")
}

func Test(t *testing.T) {
	require.Equal(t, true, math.IsNaN(Calculator("2 5")), "2 5 = NaN")
	require.Equal(t, true, math.IsNaN(Calculator("2_5")), "2_5 = NaN")
	require.Equal(t, 9., Calculator("(1+2)*3"), "(1+2)*3 = 9")
	require.Equal(t, -5., Calculator("1+2*(-3)"), "1+2*(-3) = 7")
	require.Equal(t, -1., Calculator("(1+2)-4"), "(1+2)-4 = -1")
	require.Equal(t, 16., Calculator("(2+2)*4"), "(2+2)*4 = 16")
	require.Equal(t, 10., Calculator("2+2*4"), "2+2*4 = 10")
	require.Equal(t, 0., Calculator("(2-2)*4"), "(2-2)*4 = 0")
	require.Equal(t, 0., Calculator("(2-2)*4/4"), "(2-2)*4/4 = 0")
	require.Equal(t, 345.6813704714901, Calculator("1.43*(-2.0553/9.523)+99.99+123*2"), "1.43*(-2.0553/9.523)+99.99+123*2=345.68")
	require.Equal(t, 347.6853704714901, Calculator("1.43*(-2.0553/9.523)+99.99+123*2+2.004"), "1.43*(-2.0553/9.523)+99.99+123*2+2.004=347.69")
}
