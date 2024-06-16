package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadian(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

/*
--- FAIL:
	TestSecondHandPoint/00:00:30 (0.00s)
	clockface_test.go:59: Wanted {0 -1} Point, but got {1.2246467991473515e-16 -1}
--- FAIL:
	TestSecondHandPoint/00:00:45 (0.00s)
  clockface_test.go:59: Wanted {-1 0} Point, but got {-1 -1.8369701987210272e-16}

One option to increase the accuracy of these angles would be to
use the rational type Rat from the math/big package.
But given the objective is to draw an SVG and not land on the moon,
I think we can live with a bit of fuzziness.
*/

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
