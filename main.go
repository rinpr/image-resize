package main

import (
	"fmt"
	"image-resize/img"
)

func main() {
	fmt.Println("Hello World!")
}

func resizeSingleImage(file string) {
	img.Resize(file)
}

func resizeMultipleImage() {
	file, err := img.GetImageFile()
	if err != nil {
		return
	}
	for _, filename := range file {
		img.Resize(filename)
	}
}
