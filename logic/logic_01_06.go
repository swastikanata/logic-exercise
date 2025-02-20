package logic

import slice "github.com/swastikanata/go-print-slice"

func Logic0106(n int) {
	array := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		array[n-i-1] = 3 * (i + 1)
	}

	slice.Print1DSlice(array)
}
