package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArr1(t *testing.T) {
	input := []int{1, 3, -1}
	expected := 5
	result := maxArr(input)

	assert.Equal(t, expected, result)
}

func TestMaxArr2(t *testing.T) {
	input := []int{1, 2, 3}
	expected := 4
	result := maxArr(input)

	assert.Equal(t, expected, result)
}

func TestMaxArr3(t *testing.T) {
	input := []int{1, 4, 3, -1}
	expected := 7
	result := maxArr(input)

	assert.Equal(t, expected, result)
}

func TestMaxArr4(t *testing.T) {
	input := []int{-99, -96, -89, -56, -74, -74, 41, -91, 61, -2}
	expected := 168
	result := maxArr(input)

	assert.Equal(t, expected, result)
}
