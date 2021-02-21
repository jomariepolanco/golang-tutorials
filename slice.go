/**
	* slice holds value of the same type
*/

package main

import (
	"fmt"
)

func main() {
	//composite literal
	//x := type{values}
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)

	// slice for range
	fmt.Println(len(x)) //length of slice
	fmt.Println(x[3]) //8

	for i, v := range x {
		fmt.Println(i, v)
	}

	//slicing a slice - similar to slice in JS
	fmt.Println(x[1]) //5
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	//same as range
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

	//appending a slice
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)


	//delete a slice - must use slice (like JS)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//make
	z := make([]int, 10, 100) //type, length of slice, capacity of array
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z)) //capacity of underlying array of slice

	//multi-dimensional slice aka nested slices
	/**
		* think about it like a spreadsheet
		[][]string
	*/

	jb := []string{"James", "Bond", "Choco", "martini"}
	mp := []string{"Miss", "Moneypenny", "Stawberry", "Bubblegum"}
	
	xp := [][]string{jb, mp}
	fmt.Println(xp)
}