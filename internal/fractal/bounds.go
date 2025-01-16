package fractal

import (
	"math"
	"math/rand"
)

func (d *DiamondSquareGenerator) GenerateWithBounds(
	n int, k float64, left, right, top, bottom []float64,
) ([][]float64, int) {
	return GenerateDiamondSquareWithBoundaries(n, k, left, right, top, bottom)
}

func GenerateDiamondSquareWithBoundaries(
	n int, k float64, left, right, top, bottom []float64,
) ([][]float64, int) {
	gridSize := int(math.Pow(2, float64(n))) + 1
	grid := initializeGridWithBoundaries(gridSize, left, right, top, bottom)
	applyDiamondSquare(grid, gridSize, n, k)
	return grid, gridSize
}

func initializeGridWithBoundaries(
	gridSize int, left, right, top, bottom []float64,
) [][]float64 {
	grid := make([][]float64, gridSize)
	for i := range grid {
		grid[i] = make([]float64, gridSize)
	}

	grid[0][0] = rand.Float64()
	grid[0][gridSize-1] = rand.Float64()
	grid[gridSize-1][0] = rand.Float64()
	grid[gridSize-1][gridSize-1] = rand.Float64()

	if left != nil {
		for i := 0; i < gridSize; i++ {
			grid[i][0] = left[i]
		}
	}
	if right != nil {
		for i := 0; i < gridSize; i++ {
			grid[i][gridSize-1] = right[i]
		}
	}
	if top != nil {
		copy(grid[0], top)
	}
	if bottom != nil {
		copy(grid[gridSize-1], bottom)
	}

	return grid
}

func calcSquareWithBoundary(grid [][]float64, row, col, size int, k float64) {
	capacity := len(grid)
	scale := k * 2 * float64(size) / float64(capacity)

	if grid[row][col] != 0 {
		return
	}

	sum, count := 0.0, 0
	if row-size >= 0 {
		sum += grid[row-size][col]
		count++
	}
	if col-size >= 0 {
		sum += grid[row][col-size]
		count++
	}
	if col+size < capacity {
		sum += grid[row][col+size]
		count++
	}
	if row+size < capacity {
		sum += grid[row+size][col]
		count++
	}

	grid[row][col] = (sum / float64(count)) + scale*rand.Float64()*2 - scale
}
