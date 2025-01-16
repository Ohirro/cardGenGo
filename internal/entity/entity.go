package entity

type Chunk struct {
	Grid     [][]float64
	GridSize int
}

type BoundaryDirection int

const (
	Top BoundaryDirection = iota
	Bottom
	Left
	Right
)

func (d BoundaryDirection) String() string {
	switch d {
	case Top:
		return "Top"
	case Bottom:
		return "Bottom"
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		return "Unknown"
	}
}
