package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0202(n int) {
	fmt.Println("logic_02_02")

	matrix := utils.Create2DArray(n)

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = 2*col + 2
		}
	}

	slice.Print2DSlice(matrix)

}
