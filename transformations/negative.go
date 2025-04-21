package transformations

import (
	"image/color"
)

func Negative(params TransformationParams) {
	MAX := uint32(255)
	for y := range params.OriginalImage.Bounds().Max.Y {
		r, g, b, _ := params.OriginalImage.At(params.Line, y).RGBA()

		red, green, blue := MAX-(r/MAX), MAX-(g/MAX), MAX-(b/MAX)
		c := color.RGBA{uint8(red), uint8(green), uint8(blue), 1}
		params.TransformedImage.Set(params.Line, y, c)
	}
}
