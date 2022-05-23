package main

import (
	"math"
	"testing"
)


func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 40.0
		checkPerimeter(t, rectangle, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		want := 2 * math.Pi * circle.Radius
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		want := 100.0
		checkArea(t, rectangle, want)
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := math.Pi * math.Pow(circle.Radius, 2)
		checkArea(t, circle, want)
	})
}
