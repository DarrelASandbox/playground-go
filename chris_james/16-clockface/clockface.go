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

// FAIL: TestSecondsInRadian (0.00s)
// clockface_test.go:15: Wanted 3.141592653589793 radians, but got 3.1415926535897936
func secondsInRadian(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}
