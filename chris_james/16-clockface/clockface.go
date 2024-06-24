package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) +
		(math.Pi / (30 / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

/*
We need to move the hour hand along a little bit for both the minutes and the seconds.
One full turn is one hour for the minute hand, but for the hour hand it's twelve hours.
So we just divide the angle returned by minutesInRadians by twelve
*/
func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
