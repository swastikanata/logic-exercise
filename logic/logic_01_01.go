package logic

func Logic0101(n int) []int {
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = 2*i + 1
	}

	return array
}
