package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	require.Equal(t, 1., Sum(0, 0., 0, 1.))
	require.Equal(t, 15., Sum(1, 2, 3, 4, 5))
	require.Equal(t, true, math.IsNaN(Sum(1, 2, 3, 4., "not number")))
}
