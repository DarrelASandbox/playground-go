package main

import "fmt"

func main() {
	person()
	vehicle()
	anonStruct()
}

func person() {
	fmt.Println("\n\nperson:")
	type person struct {
		first       string
		last        string
		favFlavours []string // favorite ice cream flavours
	}

	p1 := person{first: "James", last: "Bomb", favFlavours: []string{"chocolate", "martini", "rum and coke"}}
	p2 := person{first: "Larry", last: "Fairy", favFlavours: []string{"strawberry", "vanilla", "capuccino"}}

	fmt.Println(p1.first, p1.first)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2)

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println("\n", m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func vehicle() {
	fmt.Println("\n\nvehicle:")

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

	t := truck{
		vehicle:   vehicle{doors: 2, color: "red"},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{doors: 4, color: "black"},
		luxury:  false,
	}

	fmt.Println(t)
	fmt.Println(s)
}

func anonStruct() {
	fmt.Println("\n\nanonStruct:")

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
