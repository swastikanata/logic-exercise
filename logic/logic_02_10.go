package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0210(n int) {

	matrix := utils.Create2DArray(n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row >= col {
				matrix[row][col] = 2*col + 1
			}
		}
	}

	slice.Print2DSlice(matrix)
}
