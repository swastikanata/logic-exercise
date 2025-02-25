package logic

import (
	"logic-exercise/utils"
)

func SquareBlock4(row int, col int, iteration int, matrix [][]int) {
	val := 1
	n := 2*(iteration-1) + 1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i+j >= n/2 && i-j <= n/2 && j-i <= n/2 && i+j+1 <= 3*n/2 {
				if i%2 == 0 {
					matrix[row+i][col+j] = val
				} else {
					matrix[row+i][n-j-1+col] = val
				}
				val += 2
			}
		}
	}
}

func Logic0408(n int) [][]int {
	a := 1
	b := 2
	arraySize := (2*a + (n-1)*b) * n / 2
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock4(row, col, i, matrix)
		row += 2*(i-1) + 1
		col += 2*(i-1) + 1
	}

	return matrix
}
