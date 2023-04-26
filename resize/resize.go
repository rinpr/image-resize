package resize

import (
	"fmt"
	"github.com/sunshineplan/imgconv"
	"log"
)

// Documentation for image conversion.
// https://pkg.go.dev/github.com/sunshineplan/imgconv

func PngResize(fileName string, width int) {
	// Open a test image.
	path := fmt.Sprintf("image/%s.png", fileName)
	src, err := imgconv.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: width})

	// Write the resulting image as PNG.
	err = imgconv.Save("image/result.png", mark, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
	fmt.Println("Successfully exported resized image!")
}

func JpgResize(fileName string, width int) {
	// Open a test image.
	path := fmt.Sprintf("image/%s.png", fileName)
	src, err := imgconv.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: width})

	// Write the resulting image as PNG.
	err = imgconv.Save("image/result.jpeg", mark, &imgconv.FormatOption{Format: imgconv.JPEG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
	fmt.Println("Successfully exported resized image!")
}
