package clockface

import (
	"math"
	"time"
)

// A Point represents a two-dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analog clock at time 't'
// represented by a Point.
func SecondHand(t time.Time) Point {
	p := handPositionInRadians(t)
	return Point{150, p}
}

// Provides a clock hand position in radians
func handPositionInRadians(t time.Time) float64 {
	// One second is represented on an analog clock face with 2*pi/60 radians/second.
	return math.Pi / (30 / float64(t.Second()))
}

// Determines the end point of clockhand given the time on a unit circle.
func deriveHandPoint(t time.Time) Point {
	angle := handPositionInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
