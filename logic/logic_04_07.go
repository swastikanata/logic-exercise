package logic

import (
	"logic-exercise/utils"
)

func Logic0407(n int) [][]int {
	rowSize := n * (n + 1) / 2
	a := 1
	b := 2
	colSize := (2*a + (n-1)*b) * n / 2
	matrix := utils.Create2DArray(rowSize, colSize)

	row := rowSize - 1
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock3(row, col, i, matrix)
		row -= i + 1
		col += 2*(i-1) + 1
	}

	return matrix
}
