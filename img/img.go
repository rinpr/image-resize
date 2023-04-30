package img

import (
	"fmt"
	"github.com/sunshineplan/imgconv"
	"image-resize/jsonconfig"
	"log"
	"os"
	"path/filepath"
)

// Documentation for image conversion.
// https://pkg.go.dev/github.com/sunshineplan/imgconv

func GetImageFile() ([]string, error) {
	files, err := os.ReadDir(jsonconfig.GetConfig().Directory)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}

	return filenames, nil
}

func Resize(image string) {
	if jsonconfig.GetConfig().Format == "PNG" {
		pngResize(image)
	} else if jsonconfig.GetConfig().Format == "JPG" {
		jpgResize(image)
	}
}

func pngResize(file string) {
	// Open a test image.
	path := fmt.Sprintf("%s%v", jsonconfig.GetConfig().Directory, file)
	src, err := imgconv.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: jsonconfig.GetConfig().Width})

	// Write the resulting image as PNG.
	// Get the file extension
	ext := filepath.Ext(file)
	// Remove the extension from the file name
	filename := file[0 : len(file)-len(ext)]
	path = fmt.Sprintf("result/%s.png", filename)
	err = imgconv.Save(path, mark, &imgconv.FormatOption{Format: imgconv.PNG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
	fmt.Println("Successfully exported resized image!")
}

func jpgResize(file string) {
	// Open a test image.
	path := fmt.Sprintf("%s%v", jsonconfig.GetConfig().Directory, file)
	src, err := imgconv.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: jsonconfig.GetConfig().Width})

	// Write the resulting image as PNG.
	ext := filepath.Ext(file)
	// Remove the extension from the file name
	filename := file[0 : len(file)-len(ext)]
	path = fmt.Sprintf("result/%s.jpeg", filename)
	err = imgconv.Save(path, mark, &imgconv.FormatOption{Format: imgconv.JPEG})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
	fmt.Println("Successfully exported resized image!")
}
