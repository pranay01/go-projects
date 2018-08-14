package main

import "golang.org/x/tour/pic"

// Need to run go get golang.org/x/tour/pic" before
// running this code

//Implement Pic. It should return a slice of length dy, each element of which is a slice 
//of dx 8-bit unsigned integers. When you run the program, it will display your picture, 
//interpreting the integers as grayscale (well, bluescale) values.
//The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

//(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

//(Use uint8(intValue) to convert between types.

// The output is in base64 encoding, use a base64 
// decoder to visualise this - http://freeonlinetools24.com/base64-image

// Why this doesn't need to input size of image?

// The function show fixes dx and dy as consts 256

func Pic(dx, dy int) [][]uint8 {
	pict := make([][]uint8, dy)

	for i:=0;i<dy;i++ {
		pict[i] = make([]uint8, dx)
	}

	for i:=0;i<dy;i++ {
		for j:=0;j<dx;j++ {
			pict[i][j] = uint8(i*j)
		}
	}

	return pict 
}

func main() {
	pic.Show(Pic)
}