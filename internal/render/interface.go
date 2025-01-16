package render

import (
	"image"
	"image/color"
)

type Image struct {
	Pixels [][]color.Color
	Width  int
	Height int
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

func (img *Image) At(x, y int) color.Color {
	if x < 0 || x >= img.Width || y < 0 || y >= img.Height {
		return color.RGBA{0, 0, 0, 0}
	}
	return img.Pixels[y][x]
}
