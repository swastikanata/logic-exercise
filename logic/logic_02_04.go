package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0204(n int) {

	matrix := utils.Create2DArray(n)

	count := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = count
			count = count + 3
		}
	}

	slice.Print2DSlice(matrix)

}
