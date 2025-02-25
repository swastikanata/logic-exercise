package logic

import (
	"logic-exercise/utils"
)

func Logic0403(n int) [][]int {
	arraySize := n * n
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock(row, col, i, matrix)
		SquareBlock(row, arraySize-col-i, i, matrix)
		SquareBlock(arraySize-row-i, col, i, matrix)
		SquareBlock(arraySize-row-i, arraySize-col-i, i, matrix)

		row += i
		col += i
	}

	return matrix
}
