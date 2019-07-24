package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 256
	height := width

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x <= width; x++ {
		for y := 0; y <= height; y++ {
			xor := uint8(x ^ y)

			img.Set(x, y, color.RGBA{xor, xor, xor, xor})
		}
	}

	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	png.Encode(f, img)
}
