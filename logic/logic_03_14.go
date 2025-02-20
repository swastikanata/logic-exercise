package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0314(n int) {
	fmt.Println("logic_03_14")

	matrix := utils.Create2DArray(n, n)
	count := 1
	for col := 0; col < n; col++ {
		if col%2 == 0 {
			for row := 0; row < n; row++ {
				matrix[row][col] = count
				count = count + 2
			}
		} else {
			for row := n - 1; row >= 0; row-- {
				matrix[row][col] = count
				count = count + 2
			}
		}
	}

	slice.Print2DSlice(matrix)
}
