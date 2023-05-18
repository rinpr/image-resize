package main

import (
	"fmt"
	"image-resize/http"
	"image-resize/img"
	"image-resize/model"
	"image-resize/util"
	"sync"
	"time"
)

func main() {
	//fmt.Println(img.GetImageFile())
	//resizeAllfaster()
	test()
}

func test() {
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
			success := img.Resize(filename)
			fileDir, resultDir := fmt.Sprintf("image/%s", filename), fmt.Sprintf("result/%s", filename)
			data := model.ImageData{Path: fileDir, SizeBefore: util.FileSize(fileDir), SizeAfter: util.FileSize(resultDir), IsSuccess: success}
			fmt.Println(data)
			http.SendRequest(data)
		}(filename)
	}

	wg.Wait()
}
