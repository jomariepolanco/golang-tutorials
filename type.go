package main

import (
	"fmt"
)

var a int //identifier a of type int

//create our own type
type hotdog int 

var b hotdog //identifier b of type hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//a = b 
	//not possible because TYPES ARE DIFFERENT!

	a = int(b) //assign the converted to int type b to a
	//called conversion, NOT CASTING
	//this will work
}