package color

import (
	"fmt"
	"image/color"
	"math"
)

// Color ...
type Color struct {
	R      uint8
	G      uint8
	B      uint8
	A      uint8
	Weight int
}

// NewFromRGBA converts a built in Color struct
func NewFromRGBA(c color.Color) Color {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	return Color{rgba.R, rgba.G, rgba.B, rgba.A, 1}
}

// Distance from an other color
func (c Color) Distance(c2 Color) float64 {
	return math.Sqrt(
		diffSquare(c.R, c2.R) +
			diffSquare(c.G, c2.G) +
			diffSquare(c.B, c2.B),
	)
}

func diffSquare(a uint8, b uint8) float64 {
	return math.Pow(float64(a)-float64(b), 2)
}

// Luminance calculates the perceived brightness of a color on a scale of 0-1
// https://stackoverflow.com/a/596243/1207635
// other formula: (0.2126*R + 0.7152*G + 0.0722*B)
func (c Color) Luminance() float64 {
	return (float64(c.R)*0.299 + float64(c.G)*0.587 + float64(c.B)*0.114) / 255
}

// Saturation calculates colorfulness on a scale of 0-1
func (c Color) Saturation() float64 {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	avg := (max + min) / 2

	// Achromatic
	if max == min {
		return 0
	}

	delta := max - min

	if avg > .5 {
		return delta / (2 - max - min)
	}

	return delta / (max + min)
}

// Average of two colors
func (c Color) Average(c2 Color) Color {
	sumWeight := c.Weight + c2.Weight
	return Color{
		R:      uint8((int(c.R)*c.Weight + int(c2.R)*c2.Weight) / sumWeight),
		G:      uint8((int(c.G)*c.Weight + int(c2.G)*c2.Weight) / sumWeight),
		B:      uint8((int(c.B)*c.Weight + int(c2.B)*c2.Weight) / sumWeight),
		A:      uint8((int(c.A)*c.Weight + int(c2.A)*c2.Weight) / sumWeight),
		Weight: sumWeight,
	}
}

// ToHex return a hex color string
func (c Color) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}
