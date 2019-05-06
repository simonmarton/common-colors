package calculator

import (
	"log"

	"github.com/simonmarton/common-colors/color"
)

// Config defines the parameters for the calculator to use
type Config struct {
	TransparencyTreshold uint8
	IterationCount       int8
	MinLuminance         float64
	MaxLuminance         float64
}

const defaultTransparencyTreshold uint8 = 10
const defaultIterationCount int8 = 8
const defaultMinLuminance float64 = 50
const defaultMaxLuminance float64 = 230

// Calculator can group common colors
type Calculator struct {
	config Config
}

// New Calculator instance
func New(c Config) *Calculator {
	if c.TransparencyTreshold <= 0 {
		c.TransparencyTreshold = defaultTransparencyTreshold
	}

	if c.IterationCount <= 0 {
		c.IterationCount = defaultIterationCount
	}

	if c.MinLuminance < 0 {
		c.MinLuminance = defaultMinLuminance
	}

	if c.MaxLuminance > 255 || c.MaxLuminance < 1 {
		c.MaxLuminance = defaultMaxLuminance
	}

	return &Calculator{config: c}
}

// GetCommonColors ...
func (c Calculator) GetCommonColors(colors []color.Color) []string {
	log.Printf("GetCommonColors len: %d", len(colors))

	colors = c.removeInvalidColors(colors)
	log.Printf("GetCommonColors 2 len: %d", len(colors))

	return []string{"cica", "kutya"}
}

func (c Calculator) removeInvalidColors(colors []color.Color) (result []color.Color) {
	for _, col := range colors {
		l := col.Luminance()

		// fmt.Printf("color %v ; lum: %.3f\n", col, l)

		if col.A <= c.config.TransparencyTreshold || l < c.config.MinLuminance || l > c.config.MaxLuminance {
			continue
		}

		result = append(result, col)
	}

	return result
}
