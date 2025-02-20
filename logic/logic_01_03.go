package logic

import (
	"fmt"
	slice "github.com/swastikanata/go-print-slice"
)

func Logic0103(n int) {
	fmt.Println("logic_01_03")

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = 3 * (i + 1)
	}

	slice.Print1DSlice(array)
}
