// Code for https://tour.golang.org/methods/24

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	height int
	width  int
}

func loc2color(x, y int) uint8 {
	return uint8(x * y)
	//return uint8((x+y) / 2)
	//return uint8(x^y)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	v := loc2color(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
