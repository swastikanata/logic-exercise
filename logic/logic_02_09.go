package logic

import (
	"logic-exercise/utils"
)

func Logic0209(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if (row == col) || (row+col == 8) {
				matrix[row][col] = 2*col + 1
			}
		}
	}

	return matrix
}
