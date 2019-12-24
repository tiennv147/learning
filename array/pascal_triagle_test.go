package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve1(t *testing.T) {
	input := 1
	expected := [][]int{
		{1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}

func TestSolve2(t *testing.T) {
	input := 2
	expected := [][]int{
		{1},
		{1, 1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}

func TestSolve3(t *testing.T) {
	input := 3
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}

func TestSolve4(t *testing.T) {
	input := 4
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}

func TestSolve5(t *testing.T) {
	input := 5
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}

func TestSolve6(t *testing.T) {
	input := 6
	expected := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
		{1, 5, 10, 10, 5, 1},
	}
	result := solve(input)

	assert.Equal(t, expected, result)
}
