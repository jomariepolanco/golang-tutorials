package main

import (
	"fmt"
)

func main() {
	foo()

	func(x int){
		fmt.Println("Anonymous func", "The meaning of life:", x)
	}(42)

	//func expression
	//function is also a type in Go
	f := func(x int){
		fmt.Println("func expression")
		fmt.Println(x)
	}

	f(42)
}

func foo() {
	fmt.Println("foo")
}