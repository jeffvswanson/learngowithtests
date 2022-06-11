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

// Provides the second hand of a clock's hand position in radians
func secondsInRadians(t time.Time) float64 {
	// One second is represented on an analog clock face with 2*pi/60 radians/second.
	return math.Pi / (30 / float64(t.Second()))
}

// Determines the end point of a clock's second hand given the time on a unit circle.
func deriveSecondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

// Provides the minute hand position on a clock in radians
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}
