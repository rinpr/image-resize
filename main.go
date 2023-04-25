package main

import (
	"fmt"
	"github.com/sunshineplan/imgconv"
	"image-reseize/config"
	"log"
)

func main() {
	// Open a test image.
	src, err := imgconv.Open("image/7.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: config.GetWidth()})

	// Write the resulting image as PNG.
	err = imgconv.Save("image/result.png", mark, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
	fmt.Println("Successfully exported resized image!")
}
