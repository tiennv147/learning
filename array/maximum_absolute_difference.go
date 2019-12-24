package array

import "math"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func maxArr(A []int) int {
	length := len(A)
	max1, max2 := math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for i := 0; i < length; i++ {
		max1 = int(math.Max(float64(max1), float64(A[i]+i)))
		min1 = int(math.Min(float64(min1), float64(A[i]+i)))
		max2 = int(math.Max(float64(max2), float64(A[i]-i)))
		min2 = int(math.Min(float64(min2), float64(A[i]-i)))
	}

	return int(math.Max(float64(max1-min1), float64(max2-min2)))
}

func absDiff(A []int, i int, j int) float64 {
	return math.Abs(float64(A[i])-float64(A[j])) + math.Abs(float64(i)-float64(j))
}

//func maxArr(A []int) int {
//	length := len(A)
//	max := math.MinInt32
//	for i := 0; i < length; i++ {
//		for j:= i + 1; j < length; j++ {
//			cur := int(math.Abs(float64(A[i]) - float64(A[j]))) + (j - i)
//			if cur > max {
//				max = cur
//			}
//		}
//	}
//	return max
//}

// 1 2 3 4 5
//1 1
//1 2
//1 3
//1 4
//1 5
//---
//2 2
//2 3
//2 4
//2 5
//---
//3 4
//3 5
//---
//4 4
//4 5
//---
//5 5
