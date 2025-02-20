package logic

import "logic-exercise/utils"

func Logic0213(n int) {

	matrix := utils.Create2DArray(n)

	padding := 0
	for row := 0; row <= n/2; row++ {
		for col := 0 + padding; col < n-padding; col++ {
			matrix[row][col] = 2*col + 1
			matrix[n-1-row][col] = 2*col + 1
		}
		padding++
	}

	printMatrix(matrix, n)
}
