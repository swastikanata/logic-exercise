package logic

import (
	"logic-exercise/utils"
)

func Logic0308(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	padding := n / 2
	counter := 1

	for col := 0; col <= n/2; col++ {
		if col%2 != 0 {
			for row := padding; row < n-padding; row++ {
				matrix[row][col] = counter
				otherCol := n - 1 - col
				matrix[row][otherCol] = counter
				counter = counter + 2
			}
		} else {
			for row := n - padding - 1; row >= padding; row-- {
				matrix[row][col] = counter
				otherCol := n - 1 - col
				matrix[row][otherCol] = counter
				counter = counter + 2
			}
		}
		padding--
	}

	return matrix
}
