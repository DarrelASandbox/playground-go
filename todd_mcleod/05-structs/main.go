package main

import "fmt"

type person struct {
	first string // explicitly (IdentifierList)
	last  string
	age   int
}

type secretAgent struct {
	person      // implicitly (EmbeddedField)
	ltk    bool // license to kill
}

func main() {
	// value of type person
	p1 := person{first: "James", last: "Bomb", age: 81}
	p2 := person{first: "Larry", last: "Fairy", age: 18}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.last, "is", p2.age, "years old.")

	sa1 := secretAgent{
		person: person{first: "Monny", last: "Worse", age: 5_324},
		ltk:    true,
	}

	fmt.Println(sa1.first, sa1.last, "is", sa1.age, "years old. License to kill?", sa1.ltk)

	p3 := struct {
		first string
		last  string
		age   int
	}{first: "Don", last: "Don", age: 3}
	fmt.Println("Anonymous structs:", p3)
}
