package main

import (
	"fmt"
)


func main() {

	// x := 2

	// for x > 1 {
	// 	if true {
	// 		fmt.Println("True")
	// 	}
	
	// 	if false {
	// 		fmt.Println("False")
	// 	}
	
	// 	if !true {
	// 		fmt.Println("False")
	// 	}
	
	// 	if !false {
	// 		fmt.Println("True")
	// 	}
	// }

	switch (2==3) {
	case false:
		fmt.Println("this should not print")
	case (2==4):
		fmt.Println("nope")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough //check the next case as well
	case (4==4):
		fmt.Println("prints again")
	default:
		fmt.Println("this is default")
	}


	//don't use fall through, rather use default case
}