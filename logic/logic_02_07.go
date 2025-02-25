package logic

import (
	"logic-exercise/utils"
)

func Logic0207(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	for row := 0; row < n; row++ {
		col := row
		matrix[row][col] = 2*row + 1
	}
	return matrix
}
