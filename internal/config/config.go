package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	Fractal    FractalConfig `envPrefix:"FRACTAL_"`
	Render     RenderConfig  `envPrefix:"RENDER_"`
	OutputFile string        `env:"OUTPUT_FILE" envDefault:"fractal.png"`
}

type FractalConfig struct {
	N int     `env:"N" envDefault:"8"`
	K float64 `env:"K" envDefault:"0.5"`
}

type RenderConfig struct {
	ViewAngle  float64 `env:"VIEW_ANGLE" envDefault:"100.0"`
	LightAngle float64 `env:"LIGHT_ANGLE" envDefault:"45.0"`
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
