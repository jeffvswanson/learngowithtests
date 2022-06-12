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
	return derivePointFromAngle(angle)
}

// Provides the minute hand position on a clock in radians
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

// Provides the points a clock's minute hand should be at given a time.
func deriveMinuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return derivePointFromAngle(angle)
}

// derivePointFromAngle build the end point of a hand of a clock represented with a unit circle given the angle of the hand from the 12 o'clock position
func derivePointFromAngle(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
