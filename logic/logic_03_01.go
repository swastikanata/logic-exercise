package logic

import "fmt"

func Logic0301(n int) {
	fmt.Println("logic_03_01")

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

	printMatrix(matrix, n)
}
