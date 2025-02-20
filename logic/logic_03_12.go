package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0312(n int) {
	fmt.Println("logic_03_12")

	matrix := utils.Create2DArray(n, n)

	counter := 1

	for i := 0; i < n/2; i++ {
		matrix[i][n-i-1] = counter
		matrix[n-1-i][n-1-i] = counter
		counter = counter + 2
	}

	counter = 1
	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			for col := 0; col <= n/2; col++ {
				if row >= col {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter
					counter = counter + 2
				}
			}
		} else {
			for col := n / 2; col >= 0; col-- {
				if row >= col {
					matrix[row][col] = counter
					matrix[n-1-row][col] = counter
					counter = counter + 2

				}
			}
		}
	}

	slice.Print2DSlice(matrix)
}
