package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0109(n int) {
	fmt.Println("logic_01_09")

	array := make([]int, n)

	a := 3
	b := 3
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
