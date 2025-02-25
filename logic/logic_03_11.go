package logic

import (
	"logic-exercise/utils"
)

func Logic0311(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	counter := 1

	for i := 0; i < n/2; i++ {
		matrix[i][i] = counter
		matrix[n-1-i][i] = counter
		counter = counter + 2
	}

	counter = 1
	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			for col := n - 1; col >= n/2; col-- {
				if row+col >= 8 {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter
					counter = counter + 2
				}
			}
		} else {
			for col := n / 2; col < n; col++ {
				if row+col >= 8 {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter
					counter = counter + 2
				}
			}
		}
	}

	return matrix
}
