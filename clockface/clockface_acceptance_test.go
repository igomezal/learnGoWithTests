package clockface_test

import (
	"bytes"
	"clockface"
	"encoding/xml"
	"testing"
	"time"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
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

/*
	 func TestSecondHandAtMidnight(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 - 90}
		got := clockface.SecondHand(tm)

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	}

	func TestSecondHandAt30Seconds(t *testing.T) {
		tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

		want := clockface.Point{X: 150, Y: 150 + 90}
		got := clockface.SecondHand(tm)

		if got != want {
			t.Errorf("Got %v, wanted %v", got, want)
		}
	}
*/
func TestSVGWriterAtMidnight(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(t, 0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(t, 0, 0, 30),
			Line{150, 150, 150, 240},
		},
	}

	for _, c := range cases {
		t.Run(testName(t, c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(t, c.line, svg.Line) {
				t.Errorf("Expected to find the second line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(t, 0, 0, 0),
			Line{150, 150, 150, 70},
		},
	}

	for _, c := range cases {
		t.Run(testName(t, c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(t, c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(t, 6, 0, 0),
			Line{150, 150, 150, 200},
		},
	}

	for _, c := range cases {
		t.Run(testName(t, c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			clockface.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(t, c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(t testing.TB, l Line, ls []Line) bool {
	t.Helper()
	for _, line := range ls {
		if line == l {
			return true
		}
	}
	return false
}

func simpleTime(t testing.TB, hours, minutes, seconds int) time.Time {
	t.Helper()
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(test testing.TB, t time.Time) string {
	test.Helper()
	return t.Format("15:04:05")
}
