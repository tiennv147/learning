package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray1(t *testing.T) {
	A := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	max := maxSubArray(A)
	assert.Equal(t, 6, max)
}

func TestMaxSubArray2(t *testing.T) {
	A := []int{1, 2, 6}

	max := maxSubArray(A)
	assert.Equal(t, 9, max)
}

func TestMaxSubArray3(t *testing.T) {
	A := []int{10}

	max := maxSubArray(A)
	assert.Equal(t, 10, max)
}

func TestMaxSubArray4(t *testing.T) {
	A := []int{-1}

	max := maxSubArray(A)
	assert.Equal(t, -1, max)
}
