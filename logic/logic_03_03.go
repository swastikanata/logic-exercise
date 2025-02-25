package logic

import (
	"logic-exercise/utils"
)

func Logic0303(n int) [][]int {
	matrix := utils.Create2DArray(n, n)
	count := 0

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				if row+col < n {
					count = count + 2
					matrix[row][col] = count

				}
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				if row+col < n {
					count = count + 3
					matrix[row][col] = count
				}
			}
		}
	}

	return matrix
}
