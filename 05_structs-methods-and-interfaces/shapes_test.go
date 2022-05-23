package main

import (
	"math"
	"testing"
)


func TestPerimeter(t *testing.T) {
	perimeterTests := []struct{
		shape Shape
		want float64
	}{
		{shape: Circle{Radius: 10.0}, want: 2 * math.Pi * Circle{10.0}.Radius},
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
		{shape: Triangle{Base: 3.0, Height: 4.0}, want: 12.0},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct{
		shape Shape
		want float64
	}{
		{shape: Circle{Radius: 10.0}, want: math.Pi * math.Pow(Circle{10.0}.Radius, 2)},
		{shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
