package main

import (
	"fmt"
)

//1
type p struct {
	first string 
	last string
	iceCream []string
}

//3
type vehicle struct {
	doors int 
	color string 
}

type truck struct {
	vehicle 
	fourWheel bool 
}

type sedan struct {
	vehicle 
	luxury bool 
}

func main() {
	//1
	p1 := p{
		first: "Jo",
		last: "Polanco",
		iceCream: []string{"Vanilla", "Mint", "Pistachio"},
	}

	p2 := p{
		first: "Alionka",
		last: "P", 
		iceCream: []string{"Chocolate", "Chocolate", "chocolate"},
	}

	
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.iceCream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.iceCream{
		fmt.Println(i, v)
	}

	//2

	m := map[string]p{ 
		p1.last: p1, 
		p2.last: p2,
	}

	fmt.Println(m["P"])
	fmt.Println(m["Polanco"])

	fmt.Println(m["P"].first, m["P"].last)
	for i, v := range m["P"].iceCream {
		fmt.Println(i, v)
	}
	fmt.Println(m["Polanco"].first, m["Polanco"].last)
	for i, v := range m["Polanco"].iceCream {
		fmt.Println(i, v)
	}

	//3
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "grey",
		},
		luxury: true,
	}

	fmt.Println(t, s)
	fmt.Println(t.doors, s.luxury)

	lacrosse := struct {
		name string
		players int 
	}{
		name: "Lacrosse",
		players: 9,
	}
	fmt.Println(lacrosse)

} //main