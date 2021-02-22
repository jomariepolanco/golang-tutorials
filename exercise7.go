package main

import (
	"fmt"
)

type people struct {
	first string
}

func changeMe(p *people){
	(*p).first = "You" 
}

func main() {
	x := 43
	fmt.Println(&x)

	p1 := people {
		first: "Me",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}