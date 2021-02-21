package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	//type
	fmt.Printf("%T/n", y)
	
}

/** 
	* Println - general priting
	* Sprintf - print to a string which you can assign to a variable
	* FPrint - printing to a file or web server's response
*/