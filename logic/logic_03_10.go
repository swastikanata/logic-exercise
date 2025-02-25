package logic

import (
	"logic-exercise/utils"
)

func Logic0310(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	for row := 0; row <= n/2; row++ {
		for col := 0; col <= n/2; col++ {
			if row+col >= 4 {
				value := 1 + 2*(n-1-row-col)
				matrix[row][col] = value
				matrix[n-1-row][col] = value
				matrix[row][n-1-col] = value
				matrix[n-1-row][n-1-col] = value
			}
		}
	}

	return matrix
}
