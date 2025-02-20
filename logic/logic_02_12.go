package logic

import "logic-exercise/utils"

func Logic0212(n int) {

	matrix := utils.Create2DArray(n)

	padding := 0

	for i := 0; i <= n/2; i++ {
		for j := 0 + padding; j < n-padding; j++ {
			matrix[j][i] = 2*i + 1
			matrix[j][n-1-i] = 2*(n-1-i) + 1
		}
		padding++
	}

	printMatrix(matrix, n)
}
