package color

import (
	"fmt"
	"image/color"
	"math"
	"sort"
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

// Hue degree 0-1
func (c Color) Hue() float64 {
	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))

	// Achromatic
	if max == min {
		return 0
	}

	delta := max - min
	var h float64
	switch max {
	case r:
		h = (g - b) / delta
		if g < b {
			h += 6
		}
	case g:
		h = (b-r)/delta + 2
	case b:
		h = (r-g)/delta + 4
	}

	return h / 6
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

// Average of a list of colors
func Average(colors []Color) Color {
	var sumR, sumG, sumB, sumA, sumWeight int

	for _, c := range colors {
		sumR += int(c.R) * c.Weight
		sumG += int(c.G) * c.Weight
		sumB += int(c.B) * c.Weight
		sumA += int(c.A) * c.Weight
		sumWeight += c.Weight
	}

	return Color{
		R:      uint8(sumR / sumWeight),
		G:      uint8(sumG / sumWeight),
		B:      uint8(sumB / sumWeight),
		A:      uint8(sumA / sumWeight),
		Weight: sumWeight,
	}
}

// ToHex return a hex color string
func (c Color) ToHex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// Colors ...
type Colors []Color

func (a Colors) Len() int           { return len(a) }
func (a Colors) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Colors) Less(i, j int) bool { return a[i].Weight > a[j].Weight }

// Sort ...
func Sort(colors []Color) (result []Color) {
	result = colors
	sort.Sort(Colors(result))
	return result
}

// https://en.wikipedia.org/wiki/YIQ

// Y calculates the brightness compoennt, 0-255
func (c Color) Y() float64 {
	return float64(c.R)*0.29889531 + float64(c.G)*0.58662247 + float64(c.B)*0.11448223
}

// I calculates the one of the chrominance compoennts, 0-255
func (c Color) I() float64 {
	return float64(c.R)*0.59597799 - float64(c.G)*0.27417610 - float64(c.B)*0.32180189
}

// Q calculates the one of the chrominance compoennts, 0-255
func (c Color) Q() float64 {
	return float64(c.R)*0.21147017 - float64(c.G)*0.52261711 + float64(c.B)*0.31114694
}

// YIQDistance from an other color calculatedin in the YIQ color space
func (c Color) YIQDistance(c2 Color) float64 {
	y := c.Y() - c2.Y()
	i := c.I() - c2.I()
	q := c.Q() - c2.Q()

	// (0.299, 0.596, 0.212);
	// fmt.Printf("YIQ: %.4f, %.4f,%.4f\n", y, i, q)
	// fmt.Printf("YIQ: %.4f, %.4f,%.4f\n", c.Y(), c.I(), c.Q())

	return math.Sqrt(math.Pow(y, 2)*0.5053 + math.Pow(i, 2)*0.299 + math.Pow(q, 2)*0.1957)
}
