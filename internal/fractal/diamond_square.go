package fractal

import (
	"math"
	"math/rand"
	"time"
)

func GenerateDiamondSquare(n int, k float64) ([][]float64, int) {
	rand.Seed(time.Now().UnixNano())

	gridSize := int(math.Pow(2, float64(n))) + 1
	grid := initializeGrid(gridSize)

	applyDiamondSquare(grid, gridSize, n, k)
	return grid, gridSize
}

func initializeGrid(gridSize int) [][]float64 {
	grid := make([][]float64, gridSize)
	for i := range grid {
		grid[i] = make([]float64, gridSize)
	}

	grid[0][0] = rand.Float64()
	grid[0][gridSize-1] = rand.Float64()
	grid[gridSize-1][0] = rand.Float64()
	grid[gridSize-1][gridSize-1] = rand.Float64()

	return grid
}

func applyDiamondSquare(grid [][]float64, gridSize, n int, k float64) {
	size := int(math.Pow(2, float64(n)))
	for size > 1 {
		half := size / 2

		for y := half; y < gridSize; y += size {
			for x := half; x < gridSize; x += size {
				calcDiamond(grid, y, x, half, k)
			}
		}

		for y := 0; y < gridSize; y += size {
			for x := half; x < gridSize; x += size {
				calcSquare(grid, y, x, half, k)
			}
		}
		for y := half; y < gridSize; y += size {
			for x := 0; x < gridSize; x += size {
				calcSquare(grid, y, x, half, k)
			}
		}

		size /= 2
	}
}

func calcDiamond(grid [][]float64, row, col, size int, k float64) {
	capacity := len(grid)
	if row-size < 0 || row+size >= capacity || col-size < 0 || col+size >= capacity {
		return
	}

	scale := k * 2 * float64(size) / float64(capacity)
	avg := (grid[row-size][col-size] +
		grid[row-size][col+size] +
		grid[row+size][col-size] +
		grid[row+size][col+size]) / 4.0

	grid[row][col] = avg + scale*rand.Float64()*2 - scale
}

func calcSquare(grid [][]float64, row, col, size int, k float64) {
	capacity := len(grid)
	scale := k * 2 * float64(size) / float64(capacity)

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
