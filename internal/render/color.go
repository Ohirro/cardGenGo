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

var (
	PastelPalette = Palette{
		DeepWater:     color.RGBA{100, 150, 220, 255},
		ShallowWater:  color.RGBA{200, 230, 250, 255},
		Sand:          color.RGBA{240, 230, 210, 255},
		LowPlains:     color.RGBA{220, 200, 180, 255},
		Forest:        color.RGBA{150, 210, 150, 255},
		HighPlains:    color.RGBA{170, 200, 170, 255},
		Hills:         color.RGBA{190, 180, 150, 255},
		Slopes:        color.RGBA{200, 210, 190, 255},
		HighMountains: color.RGBA{220, 240, 200, 255},
		MountainPeaks: color.RGBA{230, 230, 230, 255},
	}

	WarmPastelPalette = Palette{
		DeepWater:     color.RGBA{100, 130, 200, 255},
		ShallowWater:  color.RGBA{190, 220, 240, 255},
		Sand:          color.RGBA{240, 220, 190, 255},
		LowPlains:     color.RGBA{210, 190, 170, 255},
		Forest:        color.RGBA{160, 200, 150, 255},
		HighPlains:    color.RGBA{180, 190, 160, 255},
		Hills:         color.RGBA{200, 170, 140, 255},
		Slopes:        color.RGBA{210, 200, 180, 255},
		HighMountains: color.RGBA{230, 220, 200, 255},
		MountainPeaks: color.RGBA{240, 240, 230, 255},
	}

	CoolPastelPalette = Palette{
		DeepWater:     color.RGBA{100, 130, 200, 255},
		ShallowWater:  color.RGBA{180, 220, 250, 255},
		Sand:          color.RGBA{230, 230, 220, 255},
		LowPlains:     color.RGBA{200, 210, 200, 255},
		Forest:        color.RGBA{130, 220, 170, 255},
		HighPlains:    color.RGBA{160, 200, 180, 255},
		Hills:         color.RGBA{180, 190, 200, 255},
		Slopes:        color.RGBA{190, 210, 230, 255},
		HighMountains: color.RGBA{210, 230, 240, 255},
		MountainPeaks: color.RGBA{220, 240, 255, 255},
	}

	CoolPastelPalette2 = Palette{
		DeepWater:     color.RGBA{90, 130, 200, 255},
		ShallowWater:  color.RGBA{170, 200, 240, 255},
		Sand:          color.RGBA{220, 210, 200, 255},
		LowPlains:     color.RGBA{190, 200, 220, 255},
		Forest:        color.RGBA{150, 180, 220, 255},
		HighPlains:    color.RGBA{160, 180, 220, 255},
		Hills:         color.RGBA{180, 160, 200, 255},
		Slopes:        color.RGBA{190, 190, 230, 255},
		HighMountains: color.RGBA{210, 230, 240, 255},
		MountainPeaks: color.RGBA{220, 240, 255, 255},
	}

	WarmPalette = Palette{
		DeepWater:     color.RGBA{80, 70, 200, 255},
		ShallowWater:  color.RGBA{70, 70, 250, 255},
		Sand:          color.RGBA{220, 170, 130, 255},
		LowPlains:     color.RGBA{190, 140, 100, 255},
		Forest:        color.RGBA{160, 110, 80, 255},
		HighPlains:    color.RGBA{210, 140, 100, 255},
		Hills:         color.RGBA{230, 170, 120, 255},
		Slopes:        color.RGBA{240, 190, 150, 255},
		HighMountains: color.RGBA{250, 210, 190, 255},
		MountainPeaks: color.RGBA{255, 245, 235, 255},
	}
)

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
