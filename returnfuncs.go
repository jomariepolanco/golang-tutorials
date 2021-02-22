package main

import (
	"fmt"
)

func main() {

	x := func() int {
		return 451
	}() //CALL IT

	fmt.Println("anon func", x)

	y := bar() //call it
	fmt.Printf("%T", y)
	fmt.Println("\nfunc in a func", y())
}

//func that returns a func that returns an int
func bar() func() int {
	return func() int {
		return 451
	}
}