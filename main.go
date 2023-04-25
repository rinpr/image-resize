package main

import (
	"fmt"
	"image-reseize/config"
)

func main() {
	fmt.Println(config.GetHeight())
	fmt.Println(config.GetWidth())
	fmt.Println(config.GetName())
}
