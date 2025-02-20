package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0202(n int) {

	matrix := utils.Create2DArray(n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 2*j + 2
		}
	}

	slice.Print2DSlice(matrix)

}
