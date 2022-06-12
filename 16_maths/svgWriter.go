package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	minuteHandLength = 80
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

// SVGWriter writes an SVG representation of an analog clock, showing the time t, to the writer w
func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	io.WriteString(w, svgEnd)
}

// SecondHand is the unit vector of the second hand of an analog clock at time 't'
// represented by a Point. Scales an analog clock's second hand to represent the time on the clock and the
// point the second hand should be at on the clockface as represented on an SVG image.
func secondHand(w io.Writer, t time.Time) {
	p := makeHand(deriveSecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// minuteHand is the unit vector of the minute hand on an analog clock at time 't'
// represented by a Point. Scales an analog clock's minute hand to represent the time on the clock and the
// point the minute hand should be at on the clockface as represented on an SVG image.
func minuteHand(w io.Writer, t time.Time) {
	p := makeHand(deriveMinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

// makeHand performs the scale, flip, and translate needed to convert a unit circle based point into a point
// represented on the clockface.
func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}             // scale
	p = Point{p.X, -p.Y}                              // flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY} // translate
	return p
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
