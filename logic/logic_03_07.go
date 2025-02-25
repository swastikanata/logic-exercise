package logic

import (
	"logic-exercise/utils"
)

func Logic0307(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	padding := n / 2
	counter := 1

	for row := 0; row <= n/2; row++ {
		if row%2 != 0 {
			for col := padding; col < n-padding; col++ {
				matrix[row][col] = counter
				matrix[n-1-row][col] = counter
				counter = counter + 2
			}
		} else {
			for col := n - padding - 1; col >= padding; col-- {
				matrix[row][col] = counter
				matrix[n-1-row][col] = counter
				counter = counter + 2
			}
		}
		padding--
	}

	return matrix
}
