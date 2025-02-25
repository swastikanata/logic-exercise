package logic

import (
	"logic-exercise/utils"
)

func Logic0212(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	padding := 0

	for col := 0; col <= n/2; col++ {
		for row := padding; row < n-padding; row++ {
			matrix[row][col] = 2*col + 1

			otherCol := n - 1 - col
			matrix[row][otherCol] = 2*(otherCol) + 1
		}
		padding++
	}

	return matrix
}
