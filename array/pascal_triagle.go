package array

/**
 * @input A : Integer
 *
 * @Output 2D int array.
 */
func solve(A int) [][]int {
	if A < 1 {
		return nil
	}
	result := make([][]int, A)
	for row := 0; row < A; row++ {
		col := row + 1
		result[row] = make([]int, col)
		for i := 0; i < col; i++ {
			if row == 0 || row == 1 || i == 0 || i == (col-1) {
				result[row][i] = 1
			} else {
				result[row][i] = result[row-1][i-1] + result[row-1][i]
			}
		}
	}
	return result
}
