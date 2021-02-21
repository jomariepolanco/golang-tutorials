package main

import "fmt"

func main() {
	fmt.Println("Packages Intro")
	//n, e := fmt.Println("Hello", 42, true)
	//will note work because cannot have any unused variables
	n, _ := fmt.Println("Hello", 42, true)
	//_ = ignore
	fmt.Println(n)

}

/** 
	* func Println(a...interface{}) (n int, err error)
	*will take an unlimited number of values of ANY TYPE
*/

/** 
	* variadic parameters - ...<some type> 
	* interface{} - empty interface
		* every value is also of type interface{}
		* give me as many arguments of any type as you'd like
	* _ - throw away returns
	* can't have unused variables in code
		* code pollution
		* compiler doesn't allow it
	* package.Identifier - notation
*/