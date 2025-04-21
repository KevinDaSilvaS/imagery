package files

import (
	"bytes"
	"fmt"
	"image"
	jpeg "image/jpeg"
	"os"
)

func Open(path string) (image.Image, error) {
	reader, err := os.Open(path)
	if err != nil {
		fmt.Println("opening: ", err)
		return nil, err
	}
	defer reader.Close()

	return jpeg.Decode(reader)
}

func Save(jpg *image.RGBA, saveName string) {
	var buff bytes.Buffer
	jpeg.Encode(&buff, jpg, &jpeg.Options{Quality: 100})

	os.WriteFile(saveName, buff.Bytes(), 0666)
}
