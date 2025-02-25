package logic

import (
	"logic-exercise/utils"
)

func Logic0402(n int) [][]int {
	arraySize := n * (n + 1) / 2
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := arraySize - 1
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock(row, col, i, matrix)
		row -= i + 1
		col += i
	}

	return matrix
}
