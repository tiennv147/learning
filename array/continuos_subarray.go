package array

import "math"

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
//func maxSubArray(A []int )  (int) {
//	size := len(A)
//	maxSoFar := math.MinInt32
//	maxEndingHere := 0
//	for i := 0; i < size; i++ {
//		maxEndingHere += A[i]
//		if maxSoFar < maxEndingHere {
//			maxSoFar = maxEndingHere
//		}
//		if maxEndingHere < 0 {
//			maxEndingHere = 0
//		}
//	}
//	return maxSoFar
//}

func maxSubArray(A []int) int {
	size := len(A)
	if size > 0 {
		totalMax := A[0]
		contMax := totalMax
		for i := 1; i < size; i++ {
			item := A[i]
			contMax = int(math.Max(float64(contMax+item), float64(item)))
			if contMax > totalMax {
				totalMax = contMax
			}
		}
		return totalMax
	}
	return 0
}
