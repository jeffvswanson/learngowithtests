package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// A Point represents a two-dimensional Cartesian coordinate.
type Point struct {
	X float64
	Y float64
}

// secondsInRadians provides the second hand of a clock's hand position in radians relative to 12 o'clock.
func secondsInRadians(t time.Time) float64 {
	// One second is represented on an analog clock face with 2*pi/60 radians/second.
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

// deriveSecondHandPoint determines the end point of a clock's second hand given the time on a unit circle.
func deriveSecondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	return derivePointFromAngle(angle)
}

// minutesInRadians provides the minute hand position on a clock in radians relative to 12 o'clock.
func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / secondsInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// deriveMinuteHandPoint provides the points a clock's minute hand should be at given a time.
func deriveMinuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return derivePointFromAngle(angle)
}

// hoursInRadians provides the hour hand angle on a clock in radians relative to 12 o'clock.
func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%hoursInClock)))
}

// deriveHourHandPoint provides the points a clock's hour hand should be at given a time.
func deriveHourHandPoint(t time.Time) Point {
	angle := hoursInRadians(t)
	return derivePointFromAngle(angle)
}

// derivePointFromAngle build the end point of a hand of a clock represented with a unit circle given the angle of the hand from the 12 o'clock position
func derivePointFromAngle(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
