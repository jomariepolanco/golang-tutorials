package main 

import (
	"fmt"
)

func main() {
	//1
	var array [5]int
	array[0] = 0
	array[1] = 0
	array[2] = 0
	array[3] = 0
	array[4] = 0

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)

	//2
	slice := []int{42,43,44,45,46,47,48,49,50,51}

	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", slice)

	//3
	fmt.Println(slice[:5])
	fmt.Println(slice[5:])
	fmt.Println(slice[2:7])
	fmt.Println(slice[1:6])

	//4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

	//5
	z := append(x[:3], x[6:10]...)
	fmt.Println(z)

	//6
	america := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`, }
	fmt.Println("length",len(america))
	fmt.Println("capacity", cap(america))
	for j := 0; j < len(america); j++ {
		fmt.Println(j, america[j]) 
	}

	//7
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	slice2 := [][]string{jb, mp}

	for i, v := range slice2 {
		fmt.Println(i, v)
		for j, t := range v {
			fmt.Println(j, t)
		}
	}

	//8

	mapping := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, "Martinis", "Women"},

		`moneypenny_miss` : []string{"James Bond", "Literature", "Computer Science"},

		`no_dr`: []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(mapping)

	for k, v := range mapping {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	//9
	mapping["new_record"] = []string{"Hello", "Goodbye", "Test"}
	for k, v := range mapping {
		fmt.Println(k, v)
	}

	//10
	delete(mapping, "new_record")
	for k, v := range mapping {
		fmt.Println("delete", k, v)
	}

} //main