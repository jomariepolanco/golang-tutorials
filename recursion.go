package main

import (
	"fmt"
)

func main() {
	n := factorial(4)
	fmt.Println(n)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 4 * (3 * (2 * (1)))