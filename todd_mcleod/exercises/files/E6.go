package files

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("The area of square is", s.area())
	case circle:
		fmt.Println("The area of circle is", s.area())
	}
}

func E6() {
	fmt.Println("\n\n##################################################")
	fmt.Println("E6:")

	c := circle{
		radius: 4,
	}

	s := square{
		length: 2.5,
	}

	printArea(c)
	printArea(s)
}
