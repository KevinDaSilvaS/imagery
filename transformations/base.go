package transformations

import "image"

type TransformationOptions struct {
	MaskSize   uint8
	Multiplier uint8
}

type TransformationParams struct {
	OriginalImage    image.Image
	TransformedImage *image.RGBA
	Line             int
	Opts             TransformationOptions
}
