package main 

import (
	"fmt"
)

func main() {
	//1 
	x := 42
	fmt.Println("#1")
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%#x\n", x)

	//2
	fmt.Println("#2")
	g := x == 42
	h := x <= 50
	i := x >= 32
	j := x != 41
	k := x < 12
	l := x > 12
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	//3
	fmt.Println("#3")
	const a = 42
	const b int = 42
	fmt.Println(a, "\t", b)

	//4 
	fmt.Println("#4")
	c := 42
	fmt.Printf("%d\t", c)
	fmt.Printf("%b\t", c)
	fmt.Printf("%#X\t", c)
	d := c << 1
	fmt.Printf("%d\t", d)
	fmt.Printf("%b\t", d)
	fmt.Printf("%#X\t", d)

	//5
	fmt.Println("#5")
	s := `Hello, this is a raw string literal . `
	fmt.Println(s)

	//6
	fmt.Println("#6")
	const (
		y1 = iota + 1
		y2 = iota + 1
		y3 = iota + 1
		y4 = iota + 1
	)
	fmt.Println(y1, y2, y3, y4)
}