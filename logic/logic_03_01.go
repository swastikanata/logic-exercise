package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0301(n int) {
	fmt.Println("logic_03_01")

	matrix := utils.Create2DArray(n, n)
	count := 1

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				if row >= col {
					matrix[row][col] = count
					count = count + 2
				}
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				if row >= col {
					matrix[row][col] = count
					count = count + 2
				}
			}
		}
	}

	slice.Print2DSlice(matrix)
}
