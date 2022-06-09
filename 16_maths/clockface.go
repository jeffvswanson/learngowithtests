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
