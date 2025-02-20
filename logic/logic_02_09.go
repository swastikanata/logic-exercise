package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0209(n int) {
	fmt.Println("logic_02_09")

	matrix := utils.Create2DArray(n, n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if (row == col) || (row+col == 8) {
				matrix[row][col] = 2*col + 1
			}
		}
	}

	slice.Print2DSlice(matrix)
}
