package logic

import (
	"fmt"
	"math"
)

func Logic0110(n int) {
	array := make([]int, n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			array[i] = int(math.Pow(2, float64(i/2)+1))
		}
	}

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Printf("%4d", array[i])
		} else {
			fmt.Print(" fizz")
		}
	}
}
