package main

import (
	"fmt"
)

func main() {
	x, y, z := 42, "James Bond", true

	fmt.Println("x, y, z are: ", x, y, z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

	// Zero values
	var q int
	var w string
	var e bool

	fmt.Println("\nq, w, e are: ", q, w, e)

	type chicken int
	var c chicken
	var ci int

	fmt.Println("\nc is: ", c)
	fmt.Printf("%T ", c)

	c = 42
	fmt.Println(c)

	ci = int(c)
	fmt.Println("ci is: ", ci)
	fmt.Printf("%T\n", ci)
}
