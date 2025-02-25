package logic

import (
	"logic-exercise/utils"
)

func SquareBlock3(row int, col int, n int, matrix [][]int) {
	val := 1
	padding := 0
	nCol := 2*(n-1) + 1

	for i := 0; i < n; i++ {
		if i%n != 0 {
			for j := nCol - padding - 1; j >= padding; j-- {
				matrix[row+i][col+j] = val
				val += 2
			}
			padding++
		} else {
			for j := padding; j < nCol-padding; j++ {
				matrix[row+i][col+j] = val
				val += 2
			}
			padding++
		}
	}
}

func Logic0406(n int) [][]int {
	rowSize := n * (n + 1) / 2
	a := 1
	b := 2
	colSize := (2*a + (n-1)*b) * n / 2
	matrix := utils.Create2DArray(rowSize, colSize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock3(row, col, i, matrix)
		row += i
		col += 2*(i-1) + 1
	}

	return matrix
}
