package logic

import (
	"fmt"
	"strings"
)

func Logic0209(n int) {
	cell_width := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (i == j) || (i+j == 8) {
				fmt.Printf("%*d ", cell_width, 2*j+1)
			} else {
				fmt.Print(strings.Repeat(" ", cell_width+2))
			}
		}
		fmt.Println()
	}
}
