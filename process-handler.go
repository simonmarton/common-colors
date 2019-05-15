package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math"

	"github.com/nfnt/resize"
	"github.com/simonmarton/common-colors/calculator"
	"github.com/simonmarton/common-colors/color"
	"github.com/simonmarton/common-colors/models"
	"github.com/simonmarton/common-colors/server"
)

// ProcessHandler ...
type ProcessHandler struct {
	calculator *calculator.Calculator
}

// ProcessImage ...
func (h ProcessHandler) ProcessImage(file io.Reader, imageType string, config models.CalculatorConfig) (result server.CommonColorsResp, err error) {
	fmt.Printf("Processing image with config %+v\n", config)

	h.calculator = calculator.New(config)

	img, err := openImage(file, imageType)
	if err != nil {
		return server.CommonColorsResp{}, err
	}

	img = resizeImage(img, 64, 64)
	colors := colorsFromImage(img)

	colors = h.calculator.GetCommonColors(colors)
	mainColor := colors[0]
	for _, c := range colors {
		result.Colors = append(result.Colors, server.ColorResp{
			Value:       c.ToHex(),
			Weight:      c.Weight,
			HueDistance: math.Abs(mainColor.Hue() - c.Hue()),
		})
	}

	result.Gradient = h.calculator.GenrateGradientColors(colors)

	return result, nil
}

func resizeImage(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()

	if bounds.Dx() > width || bounds.Dy() > height {
		return resize.Resize(uint(width), uint(height), img, resize.Lanczos3)
	}

	return img
}

func openImage(file io.Reader, imageType string) (image.Image, error) {
	var img image.Image
	var err error

	switch imageType {
	case "image/jpg":
		fallthrough
	case "image/jpeg":
		img, err = jpeg.Decode(file)
		if err != nil {
			return nil, err
		}
	case "image/png":
		img, err = png.Decode(file)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Not supported image format: %s", imageType)
	}

	return img, nil
}

func colorsFromImage(i image.Image) (result []color.Color) {
	b := i.Bounds()
	s := b.Size()

	for x := 0; x < s.X; x++ {
		for y := 0; y < s.Y; y++ {
			c := color.NewFromRGBA(i.At(x, y))

			result = append(result, c)
		}
	}

	return result
}

// func getColorCoords(image image.Image) (result []ColorCoord) {
// 	var transparencyTreshold uint8 = 10

// 	bounds := image.Bounds()
// 	size := bounds.Size()

// 	for x := 0; x < size.X; x++ {
// 		for y := 0; y < size.Y; y++ {
// 			c := image.At(x, y)
// 			coord := ColorCoord{RGBA: color.RGBAModel.Convert(c).(color.RGBA), Weight: 1}

// 			if coord.A < transparencyTreshold {
// 				continue
// 			}

// 			result = append(result, coord)
// 		}
// 	}

// 	return result
// }
