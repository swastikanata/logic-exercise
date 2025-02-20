package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0313(n int) {
	fmt.Println("logic_03_13")

	matrix := utils.Create2DArray(n, n)

	for row := 0; row <= n/2; row++ {
		for col := 0; col <= n/2; col++ {
			if ((row+col)%2 == 0) && (row+col >= 4) {
				matrix[row][col] = 2*col + 1
				matrix[n-1-row][col] = 2*col + 1
				matrix[row][n-1-col] = 2*col + 1
				matrix[n-1-row][n-1-col] = 2*col + 1
			}
		}
	}

	slice.Print2DSlice(matrix)
}
