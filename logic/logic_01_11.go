package logic

import (
	"fmt"
)

func Logic0111(n int) {
	array := make([]int, n/2)

	for i := 0; i < n/2; i++ {
		if i == 0 {
			array[i] = 1
		} else if i == 1 {
			array[i] = 3
		} else {
			array[i] = array[i-1] * 2
		}
	}

	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Printf("%4d", array[i/2])
		} else {
			fmt.Print(" buzz")
		}
	}
}
