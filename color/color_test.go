package color

import (
	"math"
	"testing"
)

func TestColort(t *testing.T) {
	var colorTests = []struct {
		c   Color
		lum float64
		sat float64
	}{
		{Color{R: 0, G: 0, B: 0}, 0, 0},
		{Color{R: 255, G: 255, B: 255}, 1, 0},
		{Color{R: 255, G: 0, B: 0}, .299, 1},
		{Color{R: 0, G: 255, B: 0}, .587, 1},
		{Color{R: 0, G: 0, B: 255}, .114, 1},
		{Color{R: 179, G: 168, B: 151}, .6641, .1556},
		{Color{R: 100, G: 50, B: 120}, 0.286, 0.412},
	}
	tolerance := .001

	for _, tt := range colorTests {
		l := tt.c.Luminance()
		s := tt.c.Saturation()

		if math.Abs(tt.lum-l) > tolerance {
			t.Errorf("Luminance error, expected %.4f, got %.4f", tt.lum, l)
		}

		if math.Abs(tt.sat-s) > tolerance {
			t.Errorf("Saturation error, expected %.4f, got %.4f", tt.sat, s)
		}
	}
}

func TestAverage(t *testing.T) {
	c1 := Color{R: 200, G: 100, B: 20, A: 100, Weight: 3}
	c2 := Color{R: 100, G: 25, B: 80, A: 255, Weight: 2}

	expected := Color{R: 160, G: 70, B: 44, A: 162, Weight: 5}
	got := c1.Average(c2)
	if got != expected {
		t.Errorf("Average error, expected %v, got %v", expected, got)
	}
}

func TestToHex(t *testing.T) {
	c := Color{R: 255, G: 0, B: 255}
	expected := "#ff00ff"
	got := c.ToHex()

	if got != expected {
		t.Errorf("ToHex error, expected %s, got %s", expected, got)
	}
}
