package main

import "golang.org/x/tour/pic"

import "image"
import "image/color"

type Image struct{
	// just need to implement the required functions
	// specified in the interface and the ShowImage 
	// function will work for our custom defined type
	actual_color color.Color
	boundary image.Rectangle
	col color.Model
}

func (im Image) ColorModel() color.Model{
	im.col = color.RGBAModel
	return im.col 
}
func (im Image) Bounds() image.Rectangle{
	im.boundary = image.Rect(0, 0, 256, 256)
	return im.boundary 

}

func (im Image) At(x, y int) color.Color{

	im.actual_color =  color.RGBA{uint8(2^x), uint8(2^y), uint8(x+y), 255}
	return im.actual_color 

}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
