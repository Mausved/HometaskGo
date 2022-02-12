package unique

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestXlessY(t *testing.T) {
	x := 3
	y := 5
	expected := x
	assert.Equal(t, expected, min(x, y))
}

func TestYlessX(t *testing.T) {
	x := 30
	y := 5
	expected := y
	assert.Equal(t, expected, min(x, y))
}

func TestYequalX(t *testing.T) {
	x := 5
	y := 5
	expected := x
	assert.Equal(t, expected, min(x, y))
}
