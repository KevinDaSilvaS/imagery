package main

import (
	"fmt"
	"imagery/concurrent"
	"imagery/files"
	"imagery/transformations"
)

func main() {
	fmt.Println("Imagery")
	m, err := files.Open("C:/Users/Usuario/Documents/golang/imagery/examples/fons-heijnsbroek.jpg")
	if err != nil {
		return
	}
	img := concurrent.Run(m, transformations.TransformationOptions{MaskSize: 3, Multiplier: 1}, transformations.Negative)
	files.Save(img)
}
