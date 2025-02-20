package logic

import "logic-exercise/utils"

func Logic0212(n int) {

	matrix := utils.Create2DArray(n)

	padding := 0

	for col := 0; col <= n/2; col++ {
		for row := 0 + padding; row < n-padding; row++ {
			matrix[row][col] = 2*col + 1
			matrix[row][n-1-col] = 2*(n-1-col) + 1
		}
		padding++
	}

	printMatrix(matrix, n)
}
