package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		angle     float64
	}{
		{"0 seconds", simpleTime(0, 0, 0), 0},
		{"30 seconds", simpleTime(0, 0, 30), math.Pi},
		{"45 seconds", simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{"7 seconds", simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			got := secondsInRadians(tc.time)
			if got != tc.angle {
				t.Fatalf("got %v radians, want %v", got, tc.angle)
			}
		})
	}
}

func TestDeriveSecondHandPoint(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		point     Point
	}{
		{"30 seconds", simpleTime(0, 0, 30), Point{0, -1}},
		{"45 seconds", simpleTime(0, 0, 45), Point{-1, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			got := deriveSecondHandPoint(tc.time)
			if !pointEquality(got, tc.point) {
				t.Fatalf("want %v Point, got %v", tc.point, got)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		angle     float64
	}{
		{"half hour", simpleTime(0, 30, 0), math.Pi},
		{"7 seconds", simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			got := minutesInRadians(tc.time)
			if got != tc.angle {
				t.Fatalf("want %v radians, got %v", tc.angle, got)
			}
		})
	}
}

func TestDeriveMinuteHandPoint(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		point     Point
	}{
		{"half hour", simpleTime(0, 30, 0), Point{0, -1}},
		{"45 minute mark", simpleTime(0, 45, 0), Point{-1, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			got := deriveMinuteHandPoint(tc.time)
			if !pointEquality(got, tc.point) {
				t.Fatalf("want %v Point, got %v", tc.point, got)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

// Helper function to test for float equality.
func floatEquality(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

// Helper function to test for clock Point equality due to float imprecision.
func pointEquality(a, b Point) bool {
	return floatEquality(a.X, b.X) && floatEquality(a.Y, b.Y)
}
