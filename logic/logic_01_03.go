package logic

func Logic0103(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = 3 * (i + 1)
	}

	return array
}
