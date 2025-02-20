package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0101(n int) {
	fmt.Println("logic_01_01")

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = 2*i + 1
	}

	slice.Print1DSlice(array)
}
