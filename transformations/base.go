package transformations

import "image"

type TransformationOptions struct {
	MaskSize   int
	Multiplier uint8
}

type TransformationParams struct {
	OriginalImage    image.Image
	TransformedImage *image.RGBA
	Line             int
	Opts             TransformationOptions
}
