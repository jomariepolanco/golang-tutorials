package main 

import (
	"fmt"
	"encoding/json"
)

//Uppercase for json
type movie struct {
	First string `json:"First"` 
	Last string `json:"Last"`
	Age int `json:"Age"`
}

func main() {
	p1 := movie{
		First: "James",
		Last: "Bond", 
		Age: 32,
	}
	p2 := movie{
		First: "Miss",
		Last: "Moneypenny", 
		Age: 27,
	}

	people := []movie{p1,p2,}

	fmt.Println(people)

	// returns json encoding of value
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

	//takes json and makes it go
	//json to go

	s := string(bs)
	bs2 := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs2)

	var people2 []movie 
	//people2 := []person{}
	err = json.Unmarshal(bs2, &people2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people2)
	
}