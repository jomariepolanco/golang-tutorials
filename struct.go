/** 
	* struct is a data structure that allows us to compose values of different types
	* aggregate data type
	* object/class should feel right but we don't use that language with GO. We say we created a value of type.
*/

package main 

import (
	"fmt"
)

type person struct{
	first string 
	last string 
}

//embedded structs
type secretAgent struct {
	person //embedding person into secretAgent (inheritence)
	ltk bool 

}



func main() {
	//anonymous structs - doesn't have a name
		//good to minimize code pollution
	p3 := struct {
		first string 
		last string 
		age int 
	}{
		first: "test",
		last: "test",
		age: 32,
	}

	p1 := person{
		first: "James",
		last: "Bond",
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
	}

	sa := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	fmt.Println(p1, p2, sa, p3)
	fmt.Println(p1.first, p1.last)
}