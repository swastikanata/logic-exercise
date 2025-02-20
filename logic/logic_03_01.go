package logic

func Logic0301(n int) {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	dir := 1
	i := 0
	j := 0
	count := 1

	for i < n || j < n {
		matrix[i][j] = count
		count = count + 2

		if i == j {
			i++
			j++
			dir = dir * (-1)
		} else if j == 0 {
			i++
			dir = dir * (-1)
		} else {
			j = j + dir
		}
	}

	print_matrix(matrix, n)
}
