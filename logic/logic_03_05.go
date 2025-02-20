package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0305(n int) {
	fmt.Println("logic_03_05")

	matrix := utils.Create2DArray(n)

	counter := 1

	for row := 0; row <= n/2; row++ {
		for col := 0; col <= n/2; col++ {
			if row >= col {
				matrix[row][col] = counter
				matrix[n-1-row][col] = counter
				matrix[row][n-1-col] = counter
				matrix[n-1-row][n-1-col] = counter

				counter = counter + 2
			}
		}
	}

	slice.Print2DSlice(matrix)
}
