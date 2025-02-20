package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0208(n int) {
	fmt.Println("logic_02_08")

	matrix := utils.Create2DArray(n)

	for col := 0; col < n; col++ {
		row := n - col - 1
		matrix[row][col] = 2*col + 1
	}

	slice.Print2DSlice(matrix)
}
