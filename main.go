package main

import (
	"image-resize/img"
	"image-resize/jsonconfig"
)

func main() {
	img.Resize(jsonconfig.GetConfig(0))
}
