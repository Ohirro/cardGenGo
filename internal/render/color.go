package render

import (
	"image/color"
)

type Palette struct {
	DeepWater     color.RGBA
	ShallowWater  color.RGBA
	Sand          color.RGBA
	LowPlains     color.RGBA
	Forest        color.RGBA
	HighPlains    color.RGBA
	Hills         color.RGBA
	Slopes        color.RGBA
	HighMountains color.RGBA
	MountainPeaks color.RGBA
}

var CurrentPalette = WarmPastelPalette

func getColorAndPerspectiveForHeight(height int) (color.RGBA, float64) {
	switch {
	case height > 240:
		t := float64(height-240) / 15.0
		return applyFog(CurrentPalette.MountainPeaks, t), 0.35
	case height > 220:
		t := float64(height-220) / 20.0
		return applyFog(CurrentPalette.HighMountains, t), 0.3
	case height > 200:
		t := float64(height-200) / 20.0
		return applyFog(CurrentPalette.Slopes, t), 0.28
	case height > 180:
		t := float64(height-180) / 20.0
		return applyFog(CurrentPalette.Hills, t), 0.25
	case height > 160:
		t := float64(height-160) / 20.0
		return applyFog(CurrentPalette.HighPlains, t), 0.2
	case height > 140:
		t := float64(height-140) / 20.0
		return applyFog(CurrentPalette.Forest, t), 0.15
	case height > 120:
		t := float64(height-120) / 20.0
		return applyFog(CurrentPalette.LowPlains, t), 0.12
	case height > 100:
		t := float64(height-100) / 20.0
		return applyFog(CurrentPalette.Sand, t), 0.1
	case height > 60:
		t := float64(height-60) / 40.0
		return applyFog(CurrentPalette.ShallowWater, t), 0.08
	default:
		t := float64(height) / 60.0
		return applyFog(CurrentPalette.DeepWater, t), 0.09
	}
}

func applyFog(base color.RGBA, t float64) color.RGBA {
	adjustment := int(t * 15)
	return color.RGBA{
		R: clampColor(int(base.R) + adjustment),
		G: clampColor(int(base.G) + adjustment),
		B: clampColor(int(base.B) + adjustment),
		A: base.A,
	}
}

func clampColor(value int) uint8 {
	if value < 0 {
		return 0
	}
	if value > 255 {
		return 255
	}
	return uint8(value)
}

func SetPalette(palette Palette) {
	CurrentPalette = palette
}
