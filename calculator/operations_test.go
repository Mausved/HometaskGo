package calculator

import (
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestGetOperation(t *testing.T) {
	require.Equal(t, true, math.IsNaN(getOperation(1., 2., "not correct operation")))
	require.Equal(t, 3., getOperation(1., 2., "+"))
	require.Equal(t, -1., getOperation(1., 2., "-"))
	require.Equal(t, 2., getOperation(1., 2., "*"))
	require.Equal(t, 0.5, getOperation(1., 2., "/"))
}
