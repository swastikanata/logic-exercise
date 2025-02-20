package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func SquareBlock2(row int, col int, n int, matrix [][]int) (int, int) {
	counter := 1

	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			matrix[i][j] = counter
			counter = counter + 2
		}
	}

	return row - n - 1, col + n
}

func Logic0402(n int) {
	fmt.Println("logic_04_02")

	arraySize := n * (n + 1) / 2
	matrix := utils.Create2DArray(arraySize)

	row := arraySize - 1
	col := 0

	for i := 1; i <= n; i++ {
		row, col = SquareBlock2(row, col, i, matrix)
	}

	slice.Print2DSlice(matrix)
}
