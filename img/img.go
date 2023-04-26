package img

import (
	"fmt"
	"github.com/sunshineplan/imgconv"
	"image-resize/jsonconfig"
	"log"
)

// Documentation for image conversion.
// https://pkg.go.dev/github.com/sunshineplan/imgconv

func Resize(image jsonconfig.Config) {
	if image.Format == "PNG" {
		pngResize(image.Filename, image.Width)
	} else if image.Format == "JPG" {
		jpgResize(image.Filename, image.Width)
	}
}

func pngResize(fileName string, width int) {
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

func jpgResize(fileName string, width int) {
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
