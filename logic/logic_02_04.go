package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0204(n int) {
	fmt.Println("logic_02_04")

	matrix := utils.Create2DArray(n, n)

	count := 1
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			matrix[row][col] = count
			count = count + 3
		}
	}

	slice.Print2DSlice(matrix)

}
