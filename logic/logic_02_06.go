package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0206(n int) {

	matrix := utils.Create2DArray(n)

	count := 1
	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := 0; col < n; col++ {
				matrix[row][col] = count
				count = count + 3
			}
		} else {
			for col := n - 1; col >= 0; col-- {
				matrix[row][col] = count
				count = count + 3
			}
		}
	}

	slice.Print2DSlice(matrix)

}
