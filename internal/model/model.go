package model

import (
	"cart_gen/internal/entity"
	"cart_gen/internal/fractal"
	"cart_gen/internal/render"
	"image"
	"log"
)

type FractalModel struct {
	Generator fractal.Generator
	Chunks    [][]entity.Chunk
	ChunkSize int
	n         int
	k         float64
}

func NewFractalModel(generatorName string, n int, k float64) *FractalModel {
	return &FractalModel{
		Generator: fractal.NewFractalGenerator(generatorName),
		n:         n,
		k:         k,
	}
}

func (m *FractalModel) Generate() {
	log.Println("Generating initial fractal...")
	grid, gridSize := m.Generator.Generate(m.n, m.k)
	m.Chunks = [][]entity.Chunk{
		{
			{
				Grid:     grid,
				GridSize: gridSize,
			},
		},
	}
	m.ChunkSize = gridSize
}

func getColumn(grid [][]float64, colIdx int) []float64 {
	column := make([]float64, len(grid))
	for i := range grid {
		column[i] = grid[i][colIdx]
	}
	return column
}

func (m *FractalModel) ExtendLeft() {
	log.Println("Extending fractal to the left...")

	newColumn := []entity.Chunk{}
	for i := range m.Chunks {

		var rightBoundary []float64
		if len(m.Chunks[i]) > 0 {
			rightBoundary = getColumn(m.Chunks[i][0].Grid, 0)
		}

		grid, gridSize := m.Generator.GenerateWithBounds(m.n, m.k, nil, rightBoundary, nil, nil)
		if gridSize != m.ChunkSize {
			log.Fatalf("Error: new chunk size (%d) does not match current chunk size (%d)", gridSize, m.ChunkSize)
			return
		}
		newColumn = append(newColumn, entity.Chunk{
			Grid:     grid,
			GridSize: gridSize,
		})
	}

	for i := range m.Chunks {
		m.Chunks[i] = append([]entity.Chunk{newColumn[i]}, m.Chunks[i]...)
	}
}

func (m *FractalModel) ExtendRight() {
	log.Println("Extending fractal to the right...")

	newColumn := []entity.Chunk{}
	for i := range m.Chunks {

		var leftBoundary []float64
		if len(m.Chunks[i]) > 0 {
			leftBoundary = getColumn(m.Chunks[i][len(m.Chunks[i])-1].Grid, m.ChunkSize-1)
		}

		grid, gridSize := m.Generator.GenerateWithBounds(m.n, m.k, leftBoundary, nil, nil, nil)
		if gridSize != m.ChunkSize {
			log.Fatalf("Error: new chunk size (%d) does not match current chunk size (%d)", gridSize, m.ChunkSize)
			return
		}
		newColumn = append(newColumn, entity.Chunk{
			Grid:     grid,
			GridSize: gridSize,
		})
	}

	for i := range m.Chunks {
		m.Chunks[i] = append(m.Chunks[i], newColumn[i])
	}
}

func (m *FractalModel) Render(viewAngle, lightAngle float64) (*image.RGBA, error) {
	log.Println("Rendering entire fractal...")

	rows := len(m.Chunks)
	cols := len(m.Chunks[0])

	totalSizeX := cols * m.ChunkSize
	totalSizeY := rows * m.ChunkSize

	fullGrid := make([][]float64, totalSizeY)
	for i := range fullGrid {
		fullGrid[i] = make([]float64, totalSizeX)
	}

	for rowIdx, row := range m.Chunks {
		for colIdx, chunk := range row {
			for y, gridRow := range chunk.Grid {
				for x, value := range gridRow {

					globalY := rowIdx*m.ChunkSize + y
					globalX := colIdx*m.ChunkSize + x

					if globalY >= totalSizeY || globalX >= totalSizeX {
						log.Printf("Warning: skipping out-of-bounds index (%d, %d)", globalY, globalX)
						continue
					}

					fullGrid[globalY][globalX] = value
				}
			}
		}
	}

	return render.Render(fullGrid, totalSizeX, totalSizeY, viewAngle, lightAngle), nil
}
