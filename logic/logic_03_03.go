package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0303(n int) {
	fmt.Println("logic_03_03")

	matrix := utils.Create2DArray(n)
	count := 0

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				if row+col < n {
					count = count + 2
					matrix[row][col] = count

				}
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				if row+col < n {
					count = count + 3
					matrix[row][col] = count
				}
			}
		}
	}

	slice.Print2DSlice(matrix)
}
