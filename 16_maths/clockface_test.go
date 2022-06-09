package clockface

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		condition string
		Time      time.Time
		Position  Point
	}{
		{
			"hand at midnight",
			time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC),
			Point{X: 150, Y: 150 - 90},
		},
		{
			"hand at 30 seconds",
			time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC),
			Point{X: 150, Y: 150 + 90},
		},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			got := SecondHand(tc.Time)
			if got != tc.Position {
				t.Errorf("got %v, want %v", got, tc.Position)
			}
		})
	}
}

func TestHandPositionInRadians(t *testing.T) {
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
			got := handPositionInRadians(tc.time)
			if got != tc.angle {
				t.Fatalf("got %v radians, want %v", got, tc.angle)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func TestDeriveHandPoint(t *testing.T) {
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
			got := deriveHandPoint(tc.time)
			if !pointEquality(got, tc.point) {
				t.Fatalf("want %v Point, got %v", tc.point, got)
			}
		})
	}
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
