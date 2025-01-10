package fractal

type Generator interface {
	Generate(n int, k float64) ([][]float64, int)
}

type DiamondSquareGenerator struct{}

func (d DiamondSquareGenerator) Generate(n int, k float64) ([][]float64, int) {
	return GenerateDiamondSquare(n, k)
}

func NewFractalGenerator(algorithm string) Generator {
	switch algorithm {
	case "diamond_square":
		return DiamondSquareGenerator{}
	default:
		panic("Unsupported fractal algorithm")
	}
}
