package ui

import (
	"cart_gen/internal/model"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func ShowWindow(fractalModel *model.FractalModel) {
	a := app.New()
	w := a.NewWindow("Diamond-Square Fractal")

	initialImage, err := fractalModel.Render(120, 60)
	if err != nil {
		log.Fatalf("Failed to render initial image: %v", err)
	}
	canvasImage := canvas.NewImageFromImage(initialImage)
	canvasImage.FillMode = canvas.ImageFillOriginal

	regenerateButton := widget.NewButton("Regenerate", func() {
		log.Println("Regenerating fractal...")
		fractalModel.Generate()
		updateImage(fractalModel, canvasImage)
	})

	extendLeftButton := widget.NewButton("Extend Left", func() {
		log.Println("Extending fractal to the left...")
		fractalModel.ExtendLeft()
		updateImage(fractalModel, canvasImage)
	})

	extendRightButton := widget.NewButton("Extend Right", func() {
		log.Println("Extending fractal to the right...")
		fractalModel.ExtendRight()
		updateImage(fractalModel, canvasImage)
	})

	buttons := container.NewHBox(regenerateButton, extendLeftButton, extendRightButton)
	content := container.NewVBox(canvasImage, buttons)

	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func updateImage(fractalModel *model.FractalModel, canvasImage *canvas.Image) {
	img, err := fractalModel.Render(120, 60)
	if err != nil {
		log.Printf("Error rendering image: %v", err)
		return
	}
	canvasImage.Image = img
	canvasImage.Refresh()
}
