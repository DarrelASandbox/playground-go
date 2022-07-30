package main

import "fmt"

type person struct {
	name string
}

func main() {
	x := 42
	fmt.Println("address of x ", &x)

	p1 := person{
		name: "James Bomb",
	}

	fmt.Println(p1)
	changeMe(&p1) // pass in the person address
}

func changeMe(p *person) {
	p.name = "Larry Fairy"
	fmt.Println(p)
	(*p).name = "Milo"
	fmt.Println(*p) // prints the value at the specific address
}
