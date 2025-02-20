package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0112(n int) {
	fmt.Println("logic_01_12")

	a := 1
	b := 2
	array := make([]int, n)

	for i := 0; i < n; i++ {
		array[i] = a + (i%4)*b
	}
	slice.Print1DSlice(array)
}
