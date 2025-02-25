package logic

import "logic-exercise/utils"

func Logic0409(n int) [][]int {
	a := 1
	b := 2
	arraySize := (2*a + (n-1)*b) * n / 2
	matrix := utils.Create2DArray(arraySize, arraySize)

	row := arraySize - 1
	col := 0

	for i := 1; i <= n; i++ {
		SquareBlock4(row, col, i, matrix)
		row -= 2*(i) + 1
		col += 2*(i-1) + 1
	}

	return matrix
}
