package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	width   = 6
	height  = 6
	stripes = 6

	// Note that currently the attempt to assign each stripe a unique
	// color will cause it to hang up if the number of stripes is >
	// the number of colors.

	fileName = "image.png"
)

var (
	colors = map[int]*color.RGBA{
		// RGBa: RED, GREEN, BLUE, ALPHA ch. (transparency)
		// It is just a set of colors I use in draw func
		0: {R: 100, G: 200, B: 200, A: 0xff},
		1: {R: 70, G: 70, B: 21, A: 0xff},
		2: {R: 207, G: 70, B: 110, A: 0xff},
		3: {R: 78, G: 70, B: 207, A: 0xff},
		4: {R: 207, G: 205, B: 70, A: 0xff},
		5: {R: 177, G: 37, B: 180, A: 0xff},
		6: {R: 200, G: 60, B: 200, A: 0xff},
		7: {R: 100, G: 220, B: 180, A: 0xff},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	draw(img, colors)

	// create a file
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// write an image to a file
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

func contains(slice []int, i int) bool {
	for _, v := range slice {
		if v == i {
			return true
		}
	}
	return false
}

func draw(img *image.RGBA, colors map[int]*color.RGBA) {
	var mycolors []int
	for i := 0; i < stripes; i++ {
		color := rand.Intn(8)
		for contains(mycolors, color) {
			color = rand.Intn(8)
		}
		mycolors = append(mycolors, color)
	}
	orientation := rand.Intn(4)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch orientation {
			case 0:
				img.Set(x, y, colors[mycolors[x%stripes]])
			case 1:
				img.Set(x, y, colors[mycolors[y%stripes]])
			case 2:
				img.Set(x, y, colors[mycolors[(x+y)%stripes]])
			case 3:
				img.Set(x, y, colors[mycolors[(stripes+x-y)%stripes]])
			}
		}
	}
}
