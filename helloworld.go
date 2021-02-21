package main

import "fmt"

func main() {
	//(1) sequence
	fmt.Println("Hello everyone")
	foo()


	//(2) loop; iterative
	for i := 0; i < 100; i++ {

		//(3) conditional

		if i % 2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
} //PROGRAM IS DONE WHEN YOU LEAVE MAIN FUNCTION

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("and then we exited")
}

//control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional