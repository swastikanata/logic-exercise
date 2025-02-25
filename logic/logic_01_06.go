package logic

func Logic0106(n int) []int {

	array := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		array[n-i-1] = 3 * (i + 1)
	}

	return array
}
