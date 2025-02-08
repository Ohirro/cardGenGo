package render

import "image/color"

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
