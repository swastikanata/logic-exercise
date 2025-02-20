package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0306(n int) {
	fmt.Println("logic_03_06")

	matrix := utils.Create2DArray(n, n)

	padding := 0
	counter := 1

	for row := 0; row <= n/2; row++ {
		if row%2 == 0 {
			for col := 0 + padding; col < n-padding; col++ {
				matrix[row][col] = counter

				otherRow := n - 1 - row
				matrix[otherRow][col] = counter

				counter = counter + 2
			}
		} else {
			for col := n - padding - 1; col >= padding; col-- {
				matrix[row][col] = counter

				otherRow := n - 1 - row
				matrix[otherRow][col] = counter

				counter = counter + 2
			}
		}

		padding++
	}

	slice.Print2DSlice(matrix)
}
