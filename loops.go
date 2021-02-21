package main

import (
	"fmt"
)

func main() {
	//for loop
	//for init; condition; post {}
	//for clause
	/**
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
	*/

	//nested loops

	/**
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("outer loop: %d\t inner loop: %d\n", i, j)
		}
	}
	*/

	// for statement
	// for [Condition | for clause | range clause] block

	//while statement in go
	x := 1
	/**
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
	*/

	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
}