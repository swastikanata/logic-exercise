package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0213(n int) {
	fmt.Println("logic_02_13")

	matrix := utils.Create2DArray(n)

	padding := 0
	for row := 0; row <= n/2; row++ {
		for col := 0 + padding; col < n-padding; col++ {
			matrix[row][col] = 2*col + 1

			otherRow := n - 1 - row
			matrix[otherRow][col] = 2*col + 1
		}
		padding++
	}

	slice.Print2DSlice(matrix)
}
