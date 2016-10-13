// Code for https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func loc2color(x, y int) uint8 {
	//return uint8(x * y)
	//return uint8((x+y) / 2)
	return uint8(x ^ y)
}

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	for i := dy - 1; i >= 0; i-- {
		img[i] = make([]uint8, dx)
		for j := dx - 1; j >= 0; j-- {
			img[i][j] = loc2color(i, j)
		}
	}
	return img
}

func main() {
	pic.Show(Pic)

}
