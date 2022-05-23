package main

import (
	"math"
	"testing"
)


func TestPerimeter(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Perimeter()
		want := 2 * math.Pi * circle.Radius
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := math.Pi * math.Pow(circle.Radius, 2)
		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
