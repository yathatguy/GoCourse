package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	drawRectangle()
	drawLines()
}

func drawLines() {
	red := color.RGBA{200, 30, 30, 255}
	green := color.RGBA{0, 255, 0, 255}

	img := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.Draw(img, img.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	for x := 20; x < 380; x++ {
		y := x/3 + 15
		img.Set(x, y, red)
	}

	file, err := os.Create("lines.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, img)

}

func drawRectangle() {
	green := color.RGBA{0, 255, 0, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
