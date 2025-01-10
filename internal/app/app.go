package app

import (
	"cart_gen/internal/config"
	"cart_gen/internal/fractal"
	"cart_gen/internal/render"
	"log"
	"os"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	log.Printf("Configuration loaded: %+v\n", cfg)

	fractalGen := fractal.NewFractalGenerator("diamond_square")

	log.Println("Generating fractal...")
	grid, gridSize := fractalGen.Generate(cfg.Fractal.N, cfg.Fractal.K)

	log.Println("Rendering image...")
	img := render.Render(grid, gridSize, cfg.Render.ViewAngle, cfg.Render.LightAngle)

	outputFile := cfg.OutputFile
	log.Printf("Saving image to %s\n", outputFile)
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	if err = render.SaveImage(img, file); err != nil {
		return err
	}

	log.Printf("Fractal successfully saved to %s\n", outputFile)
	return nil
}
