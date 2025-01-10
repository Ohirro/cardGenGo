package render

import (
	"image"
	"image/png"
	"io"
)

func Render(grid [][]float64, gridSize int, viewAngle, lightAngle float64) *image.RGBA {
	maxHeight := findMaxHeight(grid)
	img := image.NewRGBA(image.Rect(0, 0, gridSize, gridSize))

	for y := 0; y < gridSize; y++ {
		for x := 0; x < gridSize; x++ {
			height := grid[y][x]
			color, _ := getColorAndPerspectiveForHeight(int(height * 255 / maxHeight))
			img.Set(x, y, color)
		}
	}

	return img
}

func findMaxHeight(grid [][]float64) float64 {
	maxHeight := 0.0
	for _, row := range grid {
		for _, value := range row {
			if value > maxHeight {
				maxHeight = value
			}
		}
	}
	return maxHeight
}

func SaveImage(img *image.RGBA, w io.Writer) error {
	return png.Encode(w, img)
}
