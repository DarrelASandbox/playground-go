package shapes

import "math"

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
