package logic

import (
	"logic-exercise/utils"
)

func Logic0208(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	for col := 0; col < n; col++ {
		row := n - col - 1
		matrix[row][col] = 2*col + 1
	}
	return matrix
}
