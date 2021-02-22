//keep scope of variable as little as possible
//use code block to enclose variable to limit its scope (narrow its scope)

package main 

import (
	"fmt"
)

//package level scope
var x int 

func main() {
	//narrowed to func main
	var y int 
	fmt.Println(x, y)
	x++
	y++
	fmt.Println(x, y)
	foo()
	fmt.Println(x, y)

	a := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := incrementor()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	x++
	//y++ - won't work because y is undefined in foo func
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}