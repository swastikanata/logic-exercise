package logic

import (
	"logic-exercise/utils"
)

func Logic0205(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	count := 1
	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				matrix[row][col] = count
				count = count + 2
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				matrix[row][col] = count
				count = count + 2
			}
		}
	}

	return matrix
}
