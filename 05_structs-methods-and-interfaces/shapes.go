package main

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Perimeter() float64 {
	b := t.Base
	h := t.Height
	c := math.Sqrt(math.Pow(b, 2) + math.Pow(h, 2))
	return b + h + c
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
