package transformations

import (
	"image/color"
	"math"
)

func Blur(params TransformationParams) {
	MAX := uint32(255)
	maskSize := uint32(params.Opts.MaskSize * params.Opts.MaskSize)
	bounds := params.OriginalImage.Bounds()
	maxY := bounds.Max.Y
	maxX := bounds.Max.X
	limitBackAndForth := params.Opts.MaskSize / 2
	var x, y int
	for col := range maxY {
		var sumR, sumG, sumB uint32
		for X := params.Line - limitBackAndForth; X <= params.Line+limitBackAndForth; X++ {
			for Y := col - limitBackAndForth; Y <= col+limitBackAndForth; Y++ {
				x = int(math.Abs(float64(X)))
				if X > maxX {
					x = maxX
				}

				y = int(math.Abs(float64(Y)))
				if Y > maxY {
					y = maxY
				}

				r, g, b, _ := params.OriginalImage.At(x, y).RGBA()

				sumR += (r / MAX)
				sumG += (g / MAX)
				sumB += (b / MAX)
			}
		}

		red, green, blue := sumR/maskSize, sumG/maskSize, sumB/maskSize

		c := color.RGBA{uint8(red), uint8(green), uint8(blue), 1}
		params.TransformedImage.Set(params.Line, y, c)
	}
}
