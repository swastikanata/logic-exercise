package logic

import (
	"logic-exercise/utils"
)

func Logic0312b(n int) [][]int {
	expandedN := n + n/2
	matrix := utils.Create2DArray(n, expandedN)

	counter := 1

	for i := 0; i < n/2; i++ {
		matrix[i][expandedN-i-1] = counter
		matrix[n-1-i][expandedN-1-i] = counter
		counter = counter + 2
	}

	counter = 1
	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			for col := n / 2; col <= n; col++ {
				if row+n/2 >= col {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter

					otherCol := n/2 - (col - n/2)
					matrix[row][otherCol] = counter
					matrix[n-1-row][otherCol] = counter
					counter = counter + 2
				}
			}
		} else {
			for col := n; col >= n/2; col-- {
				if row+n/2 >= col {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter

					otherCol := n/2 - (col - n/2)
					matrix[row][otherCol] = counter
					matrix[n-1-row][otherCol] = counter
					counter = counter + 2

				}
			}
		}
	}

	return matrix
}
