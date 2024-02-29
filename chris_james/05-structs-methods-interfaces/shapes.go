package shapes

import "math"

// Interfaces are a very powerful concept in statically typed languages like Go because
// they allow you to make functions that can be used with different types and
// create highly-decoupled code whilst still maintaining type-safety.

// In Go interface resolution is implicit.
// If the type you pass in matches what the interface is asking for, it will compile.

// Notice how our helper does not need to concern itself with whether the shape is a Rectangle or a Circle or a Triangle.
// By declaring an interface, the helper is decoupled from the concrete types and only has the method it needs to do its job.

// This kind of approach of using interfaces to declare only what you need is very important in software design and will be covered in more detail in later sections.
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// The syntax for declaring methods is almost the same as functions and that's because they're so similar.
// The only difference is the syntax of the method receiver `func (receiverName ReceiverType) MethodName(args)`.

// When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable.
// In many other programming languages this is done implicitly and you access the receiver via `this`.

// It is a convention in Go to have the receiver variable be the first letter of the type.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
