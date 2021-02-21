//unordered list
//hash

package main

import (
	"fmt"
)

func main() {

	//composite literal
	m := map[string]int{
		"James": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["Barnabas"]) //will give zero value

	v, ok := m["Barnabas"]
	fmt.Println(v, ok) // 0 false <- doesn't exist
	v, ok = m["James"]
	fmt.Println(v, ok) // 32 true <- exists

	// add element
	m["todd"] = 33

	// range
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete
	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian") 
	fmt.Println(m) //ian doesn't exist, so use comma ok idiom

	if v, ok := m["Ian"]; ok {
		delete(m, "Ian")
		fmt.Println(v, ok)
	}
}