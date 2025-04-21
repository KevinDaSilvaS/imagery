package concurrent

import (
	"fmt"
	"image"
	"imagery/transformations"
	"sync"
)

func Run(originalImage image.Image) *image.RGBA {
	wg := new(sync.WaitGroup)
	wg.Add(originalImage.Bounds().Max.X) //adds a goroutine for each line on image matrix

	img := image.NewRGBA(image.Rect(0, 0, originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y))
	for i := 0; i < originalImage.Bounds().Max.X; i++ {
		opts := transformations.TransformationOptions{MaskSize: 1, Multiplier: 1}
		params := transformations.TransformationParams{
			OriginalImage:    originalImage,
			TransformedImage: img,
			Line:             i,
			Opts:             opts}
		go runner(wg, params)
	}

	wg.Wait()
	return img
}

func runner(wg *sync.WaitGroup, params transformations.TransformationParams) {
	defer wg.Done()
	fmt.Println("Goroutine:: ", params.Line)
	transformations.Negative(params)
}
