package calculator

import (
	"fmt"
	"math"

	"github.com/simonmarton/common-colors/color"
	"github.com/simonmarton/common-colors/models"
	"github.com/simonmarton/common-colors/server"
)

const defaultTransparencyTreshold uint8 = 10
const defaultIterationCount int8 = 8
const defaultMinLuminance float64 = 0
const defaultMaxLuminance float64 = 1
const defaultDistanceThreshold float64 = 50
const defaultMinSaturation float64 = .3

// Calculator can group common colors
type Calculator struct {
	config models.CalculatorConfig
}

// New Calculator instance
func New(c models.CalculatorConfig) *Calculator {
	if c.TransparencyTreshold <= 0 {
		c.TransparencyTreshold = defaultTransparencyTreshold
	}

	if c.IterationCount <= 0 {
		c.IterationCount = defaultIterationCount
	}

	if c.MinLuminance < 0 {
		c.MinLuminance = defaultMinLuminance
	}

	if c.MaxLuminance > 1 || c.MaxLuminance < .0001 {
		c.MaxLuminance = defaultMaxLuminance
	}

	// TODO write test
	if c.MinSaturation < 0 || c.MinSaturation > 1 {
		c.MinSaturation = defaultMinSaturation
	}

	// TODO write test
	// 441.67 is the max distance in a cube with 255 long sides
	if c.DistanceThreshold < 1 || c.DistanceThreshold > 441.67 {
		c.DistanceThreshold = defaultDistanceThreshold
	}

	return &Calculator{config: c}
}

// GetCommonColors ...
func (c Calculator) GetCommonColors(colors []color.Color) (result server.CommonColorsResp) {
	// log.Printf("GetCommonColors len: %d", len(colors))

	colors = c.removeInvalidColors(colors)
	// log.Printf("GetCommonColors 2 len: %d", len(colors))

	for i := int8(0); i < c.config.IterationCount; i++ {
		// colors = c.groupByNearestNeighbor(colors)

		threshold := c.config.DistanceThreshold*float64(i)/float64(c.config.IterationCount-1) + 10

		// log.Printf("GetCommonColors i: %d len: %d, threshold: %.4f", i, len(colors), threshold)
		colors = c.groupByThreshold(colors, threshold)
	}

	color.Sort(colors)

	mainColor := colors[0]

	for _, c := range colors {
		result.Colors = append(result.Colors, server.ColorResp{
			Value:       c.ToHex(),
			Weight:      c.Weight,
			HueDistance: math.Abs(mainColor.Hue() - c.Hue()),
		})
	}

	return result
}

func (c Calculator) removeInvalidColors(colors []color.Color) (result []color.Color) {
	for _, col := range colors {
		l := col.Luminance()
		s := col.Saturation()

		// todo
		if col.A <= c.config.TransparencyTreshold ||
			l < c.config.MinLuminance ||
			l > c.config.MaxLuminance ||
			s < c.config.MinSaturation {
			continue
		}

		result = append(result, col)
	}

	return result
}

func (c Calculator) groupByThreshold(colors []color.Color, threshold float64) (result []color.Color) {
	for len(colors) > 1 {
		sample := colors[0]

		similarColors := []color.Color{sample}
		remainingColors := []color.Color{}

		for _, color := range colors[1:] {
			// d := sample.Distance(color)
			d := sample.YIQDistance(color)

			if d < threshold {
				similarColors = append(similarColors, color)
			} else {
				remainingColors = append(remainingColors, color)
			}
		}

		colors = remainingColors

		result = append(result, color.Average(similarColors))
	}

	// TODO? colors at this point can still have some very different colors
	// should we append them too?
	// meh, why not?
	return append(result, colors...)
}

// Nope
func (c Calculator) groupByNearestNeighbor(colors []color.Color) (result []color.Color) {
	for len(colors) > 1 {
		sample := colors[0]
		distance := math.MaxFloat64
		closestColorIdx := -1
		for idx, color := range colors[1:] {
			d := sample.Distance(color)
			if d < distance {
				distance = d
				closestColorIdx = idx + 1
			}
		}

		// fmt.Printf("groupByNearestNeighbor closestColorIdx: %d\n", closestColorIdx)

		result = append(result, sample.Average(colors[closestColorIdx]))
		// Remove remove sample (first item) and the closes color
		colors = append(colors[1:closestColorIdx], colors[closestColorIdx+1:]...)
	}

	if len(colors) > 0 {
		fmt.Println("groupByNearestNeighbor colors remained")
		result = append(result, colors...)
	}

	return result
}
