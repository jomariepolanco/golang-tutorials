package main

import (
	"fmt"
)

func main() {
	//1
	for i := 0; i < 10001; i++ {
		fmt.Println(i)
	}

	//2
	for i := 65; i < 91; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}

	//3
	x := 2021
	for x >= 1993 {
		fmt.Println(x)
		x--
	}

	//4
	x := 2022
	for {
		if x < 1994 {
			break
		}
		if x >= 1993 {
			x--
		}
		fmt.Println(x)
	}

	//5
	for x := 10; x <= 100; x++ {
		fmt.Println(x % 4)
	}


	//6
	x := true 
	if x {
		fmt.Println("true")
		x = false
		//7
	} else if !x {
		fmt.Println("false")
	} else {
		fmt.Println("hello")
	}

	//8
	switch {
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	default:
		fmt.Println("hello")
	}

	//9
	var favSport string 
	switch favSport {
	case "Baseball":
		fmt.Println("baseball")
	case "Soccer":
		fmt.Println("soccer")
	case "Lacrosse":
		fmt.Println("lax")
	default:
		fmt.Println("that's not a sport")
	}

	//10
	/**
		* true && true => true
		* true && false => false
		* true || true => true
		* true || false => true
		* !true => false

	*/
	
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	
} //main