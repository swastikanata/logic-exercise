package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0212(n int) {
	fmt.Println("logic_02_12")

	matrix := utils.Create2DArray(n)

	padding := 0

	for col := 0; col <= n/2; col++ {
		for row := 0 + padding; row < n-padding; row++ {
			matrix[row][col] = 2*col + 1

			otherCol := n - 1 - col
			matrix[row][otherCol] = 2*(otherCol) + 1
		}
		padding++
	}

	slice.Print2DSlice(matrix)
}
