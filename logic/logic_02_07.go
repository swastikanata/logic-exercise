package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0207(n int) {
	fmt.Println("logic_02_07")

	matrix := utils.Create2DArray(n)

	for row := 0; row < n; row++ {
		col := row
		matrix[row][col] = 2*row + 1
	}

	slice.Print2DSlice(matrix)

}
