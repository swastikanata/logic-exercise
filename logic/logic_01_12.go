package logic

func Logic0112(n int) []int {
	a := 1
	b := 2
	array := make([]int, n)

	for i := 0; i < n; i++ {
		array[i] = a + (i%4)*b
	}
	return array
}
