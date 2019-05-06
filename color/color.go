package color

import (
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

/*
function rgbToHsl(r, g, b) {
  r /= 255, g /= 255, b /= 255;

  var max = Math.max(r, g, b), min = Math.min(r, g, b);
  var h, s, l = (max + min) / 2;

  if (max == min) {
    h = s = 0; // achromatic
  } else {
    var d = max - min;
    s = l > 0.5 ? d / (2 - max - min) : d / (max + min);

    switch (max) {
      case r: h = (g - b) / d + (g < b ? 6 : 0); break;
      case g: h = (b - r) / d + 2; break;
      case b: h = (r - g) / d + 4; break;
    }

    h /= 6;
  }

  return [ h, s, l ];
}
*/
