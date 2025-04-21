package files

import (
	"bytes"
	"fmt"
	jpeg "image/jpeg"
	"imagery/concurrent"
	"os"
)

func Open(path string) {
	reader, err := os.Open("C:/Users/Usuario/Documents/golang/imagery/examples/fons-heijnsbroek.jpg")
	if err != nil {
		fmt.Println("opening: ", err)
	}
	defer reader.Close()

	m, err := jpeg.Decode(reader)
	if err != nil {
		fmt.Println(err)
	}

	img := concurrent.Run(m)

	var buff bytes.Buffer
	jpeg.Encode(&buff, img, &jpeg.Options{Quality: 100})

	os.WriteFile("./test.jpeg", buff.Bytes(), 0666)
}
