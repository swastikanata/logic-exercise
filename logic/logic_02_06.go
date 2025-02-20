package logic

import (
	slice "github.com/swastikanata/go-print-slice"
	"logic-exercise/utils"
)

func Logic0206(n int) {

	matrix := utils.Create2DArray(n)

	dir := 1
	i := 0
	j := 0
	count := 1

	for i < n && j < n {
		matrix[i][j] = count
		count = count + 3

		if j == n-1 && dir == 1 {
			i++
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
