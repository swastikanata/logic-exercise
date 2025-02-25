package logic

import (
	"logic-exercise/utils"
)

func SquareBlock(row int, col int, n int, matrix [][]int) {
	counter := 1

	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			matrix[i][j] = counter
			counter = counter + 2
		}
	}
}

func Logic0401(n int) [][]int {
	arraySize := n * (n + 1) / 2
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock(row, col, i, matrix)
		row += i
		col += i
	}

	return matrix
}
