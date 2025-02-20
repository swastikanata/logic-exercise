package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func SquareBlock(row int, col int, n int, matrix [][]int) (int, int) {
	counter := 1

	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			matrix[i][j] = counter
			counter = counter + 2
		}
	}

	return row + n, col + n
}

func Logic0401(n int) {
	fmt.Println("logic_04_01")

	arraySize := n * (n + 1) / 2
	matrix := utils.Create2DArray(arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		row, col = SquareBlock(row, col, i, matrix)
	}

	slice.Print2DSlice(matrix)
}
