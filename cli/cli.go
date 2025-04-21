package cli

import (
	"fmt"
	"imagery/concurrent"
	"imagery/files"
	"imagery/transformations"
	"os"
)

func Entrypoint() {
	if len(os.Args) < 3 {
		help()
		return
	}
	transformationKey := os.Args[1]
	transformationsAvailable := transformations.TransformationsAvailable()

	transformation := transformationsAvailable[transformationKey]
	if transformation == nil {
		help()
		return
	}

	imgPath := os.Args[2]
	m, err := files.Open(imgPath)
	if err != nil {
		return
	}

	img := concurrent.Run(m, transformations.TransformationOptions{MaskSize: 11, Multiplier: 1}, transformation)
	files.Save(img, imgPath+transformationKey+"_result.jpeg")
}

func help() {

	Reset := "\033[0m"
	Green := "\033[32m"
	BrightMagenta := "\033[35;1m"
	Pink := "\033[38;5;199m"
	Lime := "\033[38;5;76m"
	LightPink := "\033[38;5;225m"

	fmt.Println(Lime, "	   Imagery Help")
	fmt.Println(Lime, "*********************************")
	fmt.Println(Green, "* Transformations available: *")
	fmt.Println(BrightMagenta, "./main ", Pink, "--blur ", LightPink, "path/image")
	fmt.Println(BrightMagenta, "./main ", Pink, "--negative ", LightPink, "path/image")

	fmt.Println("")
	fmt.Println(Lime, "*********************************", Reset)
}
