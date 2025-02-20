package logic

import (
	"fmt"
	"strings"
)

func printMatrix(matrix [][]int, n int) {
	cellWidth := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				fmt.Print(strings.Repeat(" ", cellWidth))
			} else {
				fmt.Printf("%*d", cellWidth, matrix[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println()
}
