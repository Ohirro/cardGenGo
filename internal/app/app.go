package app

import (
	"cart_gen/internal/config"
	"cart_gen/internal/model"
	"cart_gen/internal/ui"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	fractalModel := model.NewFractalModel("diamond_square", cfg.Fractal.N, cfg.Fractal.K)
	fractalModel.Generate()

	ui.ShowWindow(fractalModel)

	return nil
}
