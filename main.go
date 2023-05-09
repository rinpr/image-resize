package main

import (
	"fmt"
	"image-resize/http"
	"image-resize/img"
	"image-resize/model"
	"sync"
	"time"
)

func main() {
	fmt.Println(img.GetImageFile())
}

func test() {
	data := model.ImageData{Path: "/results/test11.png", SizeBefore: "152MB", SizeAfter: "102MB", IsSuccess: true}
	fmt.Println(data)
	http.SendRequest(data)

	fmt.Println("Hello World!")

	start := time.Now()

	// Call the function you want to time
	resizeAllfaster()

	end := time.Now()
	duration := end.Sub(start)

	fmt.Printf("Function took %v milliseconds to execute\n", duration.Milliseconds())
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

func resizeAllfaster() {
	// Concurrency video: https://www.youtube.com/watch?v=_uk9BN3a0eo

	file, err := img.GetImageFile()
	if err != nil {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(file))

	for _, filename := range file {
		go func(filename string) {
			defer wg.Done()
			img.Resize(filename)
		}(filename)
	}

	wg.Wait()
}
