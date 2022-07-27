package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// If you have method speak then you are type human
type human interface {
	speak()
}

func main() {
	// func (r receiver) identifier(parameters) (return(s)) {...}
	n := []int{1, 2, 3}
	fmt.Println("sum is", sum(n...))

	deferFunc()
	secretAgentMethod()
	anonFunc()
	funcExpression()
	returnFunc()
}

// Variadic parameters
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	s := 0
	for _, v := range x {
		s += v
	}

	return s
}

func deferFunc() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer but faster!")
	fmt.Println()
}

func (p person) speak() {
	fmt.Println("\nI am", p.first, p.last, "?")
}

func (sa secretAgent) speak() {
	fmt.Println("\nI am", sa.first, sa.last, "!")
}

func secretAgentMethod() {
	sa1 := secretAgent{person: person{first: "James", last: "Bomb"}, ltk: true}
	sa1.speak()
	humanMethod(sa1) // polymorphism
}

func humanMethod(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into the humanMethod", h.(person).first) // assertion
	case secretAgent:
		fmt.Println("I was passed into the humanMethod", h.(secretAgent).ltk)
	}
	fmt.Println("I was passed into the humanMethod", h)
}

func anonFunc() {
	fmt.Println("\n\nanonFunc:")
	func() {
		fmt.Println("Anonymous func!")
	}()

	func(x int) {
		fmt.Println("Anon", x)
	}(42)
}

func funcExpression() {
	fmt.Println("\n\nfuncExpression:")
	f := func() {
		fmt.Println("First Class Citizen")
	}

	f() // assign func to variable and call it
}

func returnFunc() {
	fmt.Println("\n\nreturnFunc:")
	addRef := add()
	x := addRef(5, 3)
	fmt.Println("5 + 3 =", x)
}

func add() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}
