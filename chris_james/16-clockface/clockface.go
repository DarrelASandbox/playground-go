package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

func secondsInRadian(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadian(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
