package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0208(n int) {

	matrix := utils.Create2DArray(n)

	for i := 0; i < n; i++ {
		matrix[n-i-1][i] = 2*i + 1
	}

	slice.Print2DSlice(matrix)
}
