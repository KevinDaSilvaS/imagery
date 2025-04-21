package main

import (
	"fmt"
	"imagery/concurrent"
	"imagery/files"
)

func main() {
	fmt.Println("Imagery")
	m, err := files.Open("C:/Users/Usuario/Documents/golang/imagery/examples/fons-heijnsbroek.jpg")
	if err != nil {
		return
	}
	img := concurrent.Run(m)
	files.Save(img)
}
