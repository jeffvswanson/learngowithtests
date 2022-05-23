package main

import (
	"math"
	"testing"
)


func TestPerimeter(t *testing.T) {
	perimeterTests := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "circle", shape: Circle{Radius: 10.0}, want: 2 * math.Pi * Circle{10.0}.Radius},
		{name: "rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
		{name: "triangle", shape: Triangle{Base: 3.0, Height: 4.0}, want: 12.0},
	}
	for _, tt := range perimeterTests {
		t.Run(tt.name, func(* testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.want {
				t.Errorf("%#v got %g, want %g", tt.shape, got, tt.want)
			}
		})
		
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct{
		name string
		shape Shape
		want float64
	}{
		{name: "circle", shape: Circle{Radius: 10.0}, want: math.Pi * math.Pow(Circle{10.0}.Radius, 2)},
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, want: 72.0},
		{name: "triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g, want %g", tt.shape, got, tt.want)
			}
		})
	}
}
