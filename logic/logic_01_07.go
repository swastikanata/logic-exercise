package logic

import slice "github.com/swastikanata/go-print-slice"

func Logic0107(n int) {
	array := make([]int, n)

	a := 1
	b := 2
	mid := n / 2

	for i := 0; i < mid; i++ {
		array[i] = a + i*b
		array[n-1-i] = a + i*b
	}

	if n%2 != 0 {
		array[mid] = a + mid*b
	}

	slice.Print1DSlice(array)
}
