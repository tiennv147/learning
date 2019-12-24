package array

import "fmt"

/**
 * @input A : 2D integer array
 *
 * @Output 2D int array.
 */
func diagonal(A [][]int) [][]int {
	rows := len(A)
	newRows := (rows-1)*2 + 1
	result := make([][]int, newRows)
	for row := 0; row < newRows; row++ {
		newCols := row + 1
		if row >= rows {
			newCols = newRows - row
		}
		result[row] = make([]int, newCols)
		for col := 0; col < newCols; col++ {
			x := col
			y := row - col
			fmt.Print("(", x, ",", y, ")")
			// ---
			result[row][col] = 1
		}
		fmt.Println("")
	}

	return result
}
