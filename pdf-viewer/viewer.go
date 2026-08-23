package main

import (
	"image"

	"fyne.io/fyne/v2/canvas"
)

const (
	minZoom     = 0.25
	maxZoom     = 4.0
	zoomStep    = 0.25
	defaultZoom = 1.0
)

type Viewer struct {
	image     *canvas.Image
	engine    *Engine
	pageIndex int
	pageCount int
	zoom      float32
}

func newViewer(engine *Engine, pageCount int) *Viewer {
	img := canvas.NewImageFromImage(nil)
	img.FillMode = canvas.ImageFillContain
	return &Viewer{
		image:     img,
		engine:    engine,
		pageIndex: 0,
		pageCount: pageCount,
		zoom:      1.0,
	}
}

func (v *Viewer) Show(p *renderedPage) {
	// asserting directly to image.Image
	if img, ok := p.Image.(image.Image); ok {
		v.image.Image = img
		v.image.Refresh()
	}
}

func (v *Viewer) render() error {
	page, err := v.engine.RenderPage(v.pageIndex, v.zoom)
	if err != nil {
		return err
	}

	v.Show(page)
	return nil
}

func (v *Viewer) Next() error {
	if v.pageIndex >= v.pageCount-1 {
		return nil
	}

	v.pageIndex++
	return v.render()
}

func (v *Viewer) Prev() error {
	if v.pageIndex <= 0 {
		return nil
	}

	v.pageIndex--
	return v.render()
}
