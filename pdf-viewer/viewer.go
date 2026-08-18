package main

import (
	"image"

	"fyne.io/fyne/v2/canvas"
)

type Viewer struct {
	image *canvas.Image
}

func newViewer() *Viewer {
	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	return &Viewer{image: img}
}

func (v *Viewer) Show(p *renderedPage) {
	// asserting directly to image.Image
	if img, ok := p.Image.(image.Image); ok {
		v.image.Image = img
		v.image.Refresh()
	}
}
