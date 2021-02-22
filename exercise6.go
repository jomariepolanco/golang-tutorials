package main 

import (
	"fmt"
	"math"
)

//4
type person struct {
	first string
	last string 
	age int 
}

func (p person) speak() {
	fmt.Println("Hi my name is", p.first, p.last, "and I am", p.age, "years old.")
}

//5 
type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius 
}

func (s square) area() float64 {	
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	a := s.area()
	fmt.Println(a)
}

func main() {

	sq := square {
		length: 15.5,
	}

	c := circle {
		radius: 13.345,
	}

	info(c)
	info(sq)
	
	//6
	func() {
		fmt.Println("anon")
	}()

	defer deferring()

	x := foo()
	fmt.Println(x)

	y, z := bar()
	fmt.Println(y, z)

	two := []int{1,2,3,4,5,6,7,8,9,}
	a := foo2(two...)
	fmt.Println(a)

	b := bar2(two)
	fmt.Println(b)

	p1 := person {
		"Jomarie",
		"Polanco", 
		27,
	}
	p1.speak()

	//7 
	f := func() {
		fmt.Println(7)
	}
	f()

	//8
	d := eight()
	fmt.Println(d())

	//9

	e := nine(eight())
	fmt.Println(e)
}

//1
func foo() int{
	return 5
}


func bar() (int, string) {
	return 5, "Hello"
}

//2
func foo2(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total 
}

func bar2(x []int) int {
	total := 0
	for _, v := range x {
		total += v 
	}
	return total 
}

//3
func deferring() {
	fmt.Println(6)
}

//8
func eight() func() int {
	return func() int{
		return 8
	}
}

//9

func nine(func() int) {
	return func() int
}