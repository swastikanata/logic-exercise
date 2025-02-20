package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0104(n int) {
	fmt.Println("logic_01_04")

	array := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		array[n-i-1] = 2*i + 1
	}

	slice.Print1DSlice(array)
}
