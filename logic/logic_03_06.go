package logic

import (
	"logic-exercise/utils"
)

func Logic0306(n int) [][]int {
	matrix := utils.Create2DArray(n, n)

	padding := 0
	counter := 1

	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			for col := 0 + padding; col < n-padding; col++ {
				matrix[row][col] = counter

				otherRow := n - 1 - row
				matrix[otherRow][col] = counter

				counter = counter + 2
			}
		} else {
			for col := n - padding - 1; col >= padding; col-- {
				matrix[row][col] = counter

				otherRow := n - 1 - row
				matrix[otherRow][col] = counter

				counter = counter + 2
			}
		}

		padding++
	}

	return matrix
}
