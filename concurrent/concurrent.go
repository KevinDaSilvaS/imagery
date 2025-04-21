package concurrent

import (
	"fmt"
	"image"
	"imagery/transformations"
	"sync"
)

func Run(originalImage image.Image, opts transformations.TransformationOptions, transformation func(params transformations.TransformationParams)) *image.RGBA {
	wg := new(sync.WaitGroup)
	wg.Add(originalImage.Bounds().Max.X) //adds a goroutine for each line on image matrix

	img := image.NewRGBA(image.Rect(0, 0, originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y))
	for i := range originalImage.Bounds().Max.X {
		params := transformations.TransformationParams{
			OriginalImage:    originalImage,
			TransformedImage: img,
			Line:             i,
			Opts:             opts}
		go runner(wg, params, transformation)
	}

	wg.Wait()
	return img
}

func runner(wg *sync.WaitGroup, params transformations.TransformationParams, transformation func(params transformations.TransformationParams)) {
	defer wg.Done()
	fmt.Println("Goroutine:: ", params.Line)
	transformation(params)
}
