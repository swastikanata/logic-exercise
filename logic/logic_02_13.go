package logic

import "logic-exercise/utils"

func Logic0213(n int) {

	matrix := utils.Create2DArray(n)

	padding := 0
	for i := 0; i <= n/2; i++ {
		for j := 0 + padding; j < n-padding; j++ {
			matrix[i][j] = 2*j + 1
			matrix[n-1-i][j] = 2*j + 1
		}
		padding++
	}

	printMatrix(matrix, n)
}
