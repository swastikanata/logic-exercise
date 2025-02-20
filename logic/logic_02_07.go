package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0207(n int) {

	matrix := utils.Create2DArray(n)

	for i := 0; i < n; i++ {
		matrix[i][i] = 2*i + 1
	}

	slice.Print2DSlice(matrix)

}
