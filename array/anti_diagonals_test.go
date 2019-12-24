package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiagonal1(t *testing.T) {
	input := [][]int{
		{1},
	}
	expected := [][]int{
		{1},
	}
	result := diagonal(input)

	assert.Equal(t, expected, result)
}

func TestDiagonal2(t *testing.T) {
	input := [][]int{
		{1, 2},
		{3, 4},
	}
	expected := [][]int{
		{1},
		{2, 3},
		{4},
	}
	result := diagonal(input)

	assert.Equal(t, expected, result)
}

func TestDiagonal3(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	expected := [][]int{
		{1},
		{2, 4},
		{3, 5, 7},
		{6, 8},
		{9},
	}
	result := diagonal(input)

	assert.Equal(t, expected, result)
}

func TestDiagonal4(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := [][]int{
		{1},
		{2, 5},
		{3, 6, 9},
		{4, 7, 10, 13},
		{8, 11, 14},
		{12, 15},
		{16},
	}
	result := diagonal(input)

	assert.Equal(t, expected, result)
}
