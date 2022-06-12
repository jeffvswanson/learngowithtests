package clockface_test

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/jeffvswanson/learngowithtests/16_maths/clockface"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}
type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}
type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		line      Line
	}{
		{"0/12 position", simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{"6 position", simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(tc.line, svg.Line) {
				t.Errorf("got %+v coordinates in the SVG output, want %+v", svg, tc.line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		condition string
		time      time.Time
		line      Line
	}{
		{"0 minute position", simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}
	for _, tc := range cases {
		t.Run(tc.condition, func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, tc.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(tc.line, svg.Line) {
				t.Errorf("got minute hand line coordinates %+v, want %+v", svg.Line, tc.line)
			}
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(1337, time.January, 1, hours, minutes, seconds, 0, time.UTC)
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}
