package logic

import (
	"logic-exercise/utils"
)

func Logic0203(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	count := 1
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = count
			count = count + 2
		}
	}

	return matrix
}
