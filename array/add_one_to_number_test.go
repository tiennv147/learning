package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne2(t *testing.T) {
	input := []int{1, 2, 9}
	expected := []int{1, 3, 0}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne3(t *testing.T) {
	input := []int{9, 9}
	expected := []int{1, 0, 0}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne4(t *testing.T) {
	input := []int{9, 9, 9}
	expected := []int{1, 0, 0, 0}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne5(t *testing.T) {
	input := []int{9, 9, 9, 9, 9}
	expected := []int{1, 0, 0, 0, 0, 0}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne6(t *testing.T) {
	input := []int{0, 3, 7, 6, 4, 0, 5, 5, 5}
	expected := []int{3, 7, 6, 4, 0, 5, 5, 6}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}

func TestPlusOne7(t *testing.T) {
	input := []int{0, 0, 0, 3, 7, 6, 4, 0, 5, 5, 5}
	expected := []int{3, 7, 6, 4, 0, 5, 5, 6}
	result := plusOne(input)

	assert.Equal(t, expected, result)
}
