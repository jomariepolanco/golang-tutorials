package main 

import (
	"fmt"
)

type person struct {
	first string 
	last string 
}

type secretAgent struct {
	person 
	ltk bool 
}

//method can get a receiver. this attaches the function to that type. any value of that type can use this method.
func (s secretAgent) speak(){
	fmt.Println("I am", s.first, s.last)
}

//interfaces allow us to define behavior
//allows us to do polymorphism
//a value can be of more than one type
// "if you have this method, you're my type"
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type){
	case person:
		fmt.Println("I was passed into bar -", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("I was passed into bar -", h.(secretAgent).first, h.(secretAgent).last)
	default:
		fmt.Println("I called human")
	}
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - person")
}
func main() {
	//value is of type secretAgent AND of value human
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(p1)
	bar(sa1) //both type human so both work in bar func call
	bar(p1)
}