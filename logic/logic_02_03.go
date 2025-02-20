package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0203(n int) {

	matrix := utils.Create2DArray(n)

	count := 1
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = count
			count = count + 2
		}
	}

	slice.Print2DSlice(matrix)

}
