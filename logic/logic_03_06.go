package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0309(n int) {
	fmt.Println("logic_03_09")

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	dir := 1
	i := 0
	j := 0
	count := 1

	for i < n && j < n {
		matrix[i][j] = count
		count = count + 2

		if i == j && dir == 1 {
			i++
			j++
			dir = dir * (-1)
		} else if j == 0 && dir == -1 {
			i++
			dir = dir * (-1)
		} else {
			j = j + dir
		}
	}
	slice.Print2DSlice(matrix)
}
