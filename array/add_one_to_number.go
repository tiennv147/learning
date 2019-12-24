package array

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func plusOne(A []int) []int {
	length := len(A)
	zeroIndex := -1
	// Clean up first
	for i := 0; i < length; i++ {
		if A[i] == 0 {
			zeroIndex = i
		} else {
			break
		}
	}
	if zeroIndex != -1 {
		A = A[zeroIndex+1 : length]
		length = len(A)
	}
	var result = make([]int, length+1, length+1)
	// Init carry = 1, like plus one
	carry := 1
	for i := length - 1; i >= 0; i-- {
		sum := A[i] + carry
		if sum > 9 {
			carry = 1
			sum = sum % 10
		} else {
			carry = 0
		}
		result[i+1] = sum
	}
	if carry > 0 {
		result[0] = carry
	} else {
		result = result[1:]
	}
	return result
}
