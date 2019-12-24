package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiverse1(t *testing.T) {
	m := 0
	w := 0

	result := diverse(m, w)

	assert.Equal(t, 0, result)
}

func TestDiverse2(t *testing.T) {
	m := 1
	w := 1

	result := diverse(m, w)

	assert.Equal(t, 0, result)
}

func TestDiverse3(t *testing.T) {
	m := 1
	w := 2

	result := diverse(m, w)

	assert.Equal(t, 2, result)
}
