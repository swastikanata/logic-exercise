package logic

import "fmt"

func Logic0108(n int) {
	fmt.Println("logic_01_08")

	a := 2
	b := 2
	mid := n / 2

	for i := 0; i < mid; i++ {
		fmt.Print(a+i*b, " ")
	}

	if n%2 != 0 {
		fmt.Print(a+mid*b, " ")
	}

	for i := mid - 1; i >= 0; i-- {
		fmt.Print(a+i*b, " ")
	}

	fmt.Println()
}
