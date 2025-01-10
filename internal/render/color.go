package render

import (
	"image/color"
)

func getColorAndPerspectiveForHeight(height int) (color.RGBA, float64) {
	switch {
	case height > 240: // Вершины гор (снег)
		t := float64(height-240) / 15.0
		return color.RGBA{uint8(lerp(230, 255, t)), uint8(lerp(230, 255, t)), uint8(lerp(230, 255, t)), 255}, 0.35
	case height > 220: // Высокие горы
		t := float64(height-220) / 20.0
		return color.RGBA{uint8(lerp(220, 240, t)), uint8(lerp(220, 240, t)), uint8(lerp(200, 220, t)), 255}, 0.3
	case height > 200: // Склоны
		t := float64(height-200) / 20.0
		return color.RGBA{uint8(lerp(200, 220, t)), uint8(lerp(210, 230, t)), uint8(lerp(190, 210, t)), 255}, 0.28
	case height > 180: // Горные холмы
		t := float64(height-180) / 20.0
		return color.RGBA{uint8(lerp(190, 210, t)), uint8(lerp(180, 200, t)), uint8(lerp(150, 190, t)), 255}, 0.25
	case height > 160: // Высокие равнины
		t := float64(height-160) / 20.0
		return color.RGBA{uint8(lerp(170, 190, t)), uint8(lerp(200, 220, t)), uint8(lerp(170, 190, t)), 255}, 0.2
	case height > 140: // Леса
		t := float64(height-140) / 20.0
		return color.RGBA{uint8(lerp(150, 170, t)), uint8(lerp(210, 230, t)), uint8(lerp(150, 170, t)), 255}, 0.15
	case height > 120: // Низкие равнины
		t := float64(height-120) / 20.0
		return color.RGBA{uint8(lerp(220, 240, t)), uint8(lerp(200, 220, t)), uint8(lerp(180, 200, t)), 255}, 0.12
	case height > 100: // Песчаные участки
		t := float64(height-100) / 20.0
		return color.RGBA{uint8(lerp(240, 255, t)), uint8(lerp(230, 245, t)), uint8(lerp(210, 230, t)), 255}, 0.1
	case height > 60: // Мелководье
		t := float64(height-60) / 40.0
		return color.RGBA{uint8(lerp(200, 230, t)), uint8(lerp(230, 250, t)), uint8(lerp(250, 255, t)), 255}, 0.08
	default: // Глубокие воды
		t := float64(height) / 60.0
		return color.RGBA{uint8(lerp(100, 170, t)), uint8(lerp(150, 200, t)), uint8(lerp(220, 240, t)), 255}, 0.05
	}
}

func lerp(v0, v1, t float64) float64 {
	return v0 + t*(v1-v0)
}
