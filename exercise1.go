package main

import (
	"fmt"
)

//2
var a int = 42
var b string = "James Bond"
var c bool = true

//4
type pillow int
var n pillow 

//5
var h int 

func main() {
	// 1
	x, y, z := 42, "James Bond", true 
	fmt.Println("hands on exercise 1")
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// 2 
	fmt.Println("hands on exericse 2", a, b, c)
	//zero values

	// 3
	fmt.Println("hands on exercise 3")
	s := fmt.Sprintf("%v %v %v", a, b, c)
	fmt.Println(s)

	//4
	fmt.Println("hands on 4")
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	n = 42
	fmt.Println(n)

	//5
	fmt.Println("hands on 5")
	h = int(n)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
}