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

/*
1. Scale it to the length of the hand
2. Flip it over the X axis to account for the SVG having an origin in the top left hand corner
3. Translate it to the right position (so that it's coming from an origin of (150,150))
*/
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scale
	p = Point{p.X, -p.Y}            // flip
	p = Point{p.X + 150, p.Y + 150} // translate
	return p
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
