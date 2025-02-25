package logic

import (
	"logic-exercise/utils"
)

func SquareBlock5(row int, col int, n int, matrix [][]int) [][]int {

	for i := 0; i <= n/2; i++ {
		for j := 0; j <= n/2; j++ {
			if i+j >= n/2 {
				value := 1 + 2*((i+j)-n/2)
				matrix[row+i][col+j] = value
				matrix[row+n-1-i][col+j] = value
				matrix[row+i][col+n-1-j] = value
				matrix[row+n-1-i][col+n-1-j] = value
			}
		}
	}
	return matrix
}

func Logic0410(n int) [][]int {
	arraySize := n*(n+1) - 1
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		currentN := 2*(i-1) + 1
		SquareBlock5(row, col, currentN, matrix)
		SquareBlock5(row, arraySize-1-col-(currentN-1), currentN, matrix)
		SquareBlock5(arraySize-1-row-(currentN-1), col, currentN, matrix)
		SquareBlock5(arraySize-1-row-(currentN-1), arraySize-1-col-(currentN-1), currentN, matrix)
		row += i
		col += i
	}

	return matrix
}
