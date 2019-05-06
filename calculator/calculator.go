package calculator

import (
	"fmt"
	"log"
	"math"
	"strconv"

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
func (c Calculator) GetCommonColors(colors []color.Color) (result []string) {
	log.Printf("GetCommonColors len: %d", len(colors))

	colors = c.removeInvalidColors(colors)
	log.Printf("GetCommonColors 2 len: %d", len(colors))

	for i := int8(0); i < c.config.IterationCount; i++ {
		colors = c.groupByNearestNeighbor(colors)
		log.Printf("GetCommonColors i %d len: %d", i, len(colors))

	}

	for _, c := range colors {
		result = append(result, c.ToHex()+" W:"+strconv.Itoa(c.Weight))
	}

	return result
}

func (c Calculator) removeInvalidColors(colors []color.Color) (result []color.Color) {
	for _, col := range colors {
		l := col.Luminance()
		s := col.Saturation()

		// fmt.Printf("color %v ; lum: %.3f\n", col, l)

		if col.A <= c.config.TransparencyTreshold || l < c.config.MinLuminance || l > c.config.MaxLuminance || s < .3 {
			continue
		}

		result = append(result, col)
	}

	return result
}

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
