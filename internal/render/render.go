package render

import (
	"image"
	_ "image/color"
	"log"
)

func Render(grid [][]float64, gridWidth, gridHeight int, viewAngle, lightAngle float64) *image.RGBA {
	maxHeight := findMaxHeight(grid)
	img := image.NewRGBA(image.Rect(0, 0, gridWidth, gridHeight))

	for y := 0; y < gridHeight; y++ {
		if y >= len(grid) {
			log.Printf("Warning: skipping row %d as it is out of bounds", y)
			continue
		}
		for x := 0; x < gridWidth; x++ {
			if x >= len(grid[y]) {
				log.Printf("Warning: skipping column %d in row %d as it is out of bounds", x, y)
				continue
			}

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
