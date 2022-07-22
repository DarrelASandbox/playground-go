package main

import (
	"fmt"
)

const e = 42     // untyped
const f int = 43 // typed

const (
	y1 = 2022 + iota
	y2 = 2022 + iota
	y3 = 2022 + iota
	y4 = 2022 + iota
)

func main() {
	num := 42 // decimal, binary, and hex
	fmt.Printf("%d\t%b\t%#x", num, num, num)

	// operators & expressions
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	fmt.Println('\n', a, b, c, d)

	// iota
	fmt.Println('\n', y1, y2, y3, y4)
}
