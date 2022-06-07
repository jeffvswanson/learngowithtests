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
	return float64(t.Second()) * (math.Pi / 30)
}
