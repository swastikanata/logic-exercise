package logic

import (
	"fmt"
	"strings"
)

func print_matrix(matrix [][]int, n int) {
	cell_width := 3
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				fmt.Print(strings.Repeat(" ", cell_width))
			} else {
				fmt.Printf("%*d", cell_width, matrix[i][j])
			}
		}
		fmt.Println()
	}
}
