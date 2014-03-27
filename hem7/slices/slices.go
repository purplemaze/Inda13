package main

/*
Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
The choice of image is up to you. Interesting functions include x^y, (x+y)/2, and x*y.
(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
(Use uint8(intValue) to convert between types.)

Author: Daniel Cserhalmi
Version 1.0
*/

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	sliceY := make([][]uint8, dy) // dy = length of slice
	for i := range sliceY {
		sliceY[i] = make([]uint8, dx) //dx = length of slice
	}

	for y, sliceX := range sliceY {
		for x := range sliceX {
			sliceX[x] = uint8(x ^ y)
		}
	}

	return sliceY
}

func main() {
	pic.Show(Pic)
}
