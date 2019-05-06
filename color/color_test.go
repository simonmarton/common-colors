package color

import (
	"math"
	"testing"
)

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

func TestColort(t *testing.T) {
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
