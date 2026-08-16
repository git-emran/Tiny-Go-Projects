package main

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Toolbar struct {
	Container *widget.Toolbar
	ZoomIn    *widget.ToolbarAction
	ZoomOut   *widget.ToolbarAction
	ZoomReset *widget.ToolbarAction
	Forward   *widget.ToolbarAction
	Back      *widget.ToolbarAction
}

func newToolbar() *Toolbar {
	t := &Toolbar{}

	t.Back = widget.NewToolbarAction(theme.NavigateBackIcon(), func() {})
	t.Forward = widget.NewToolbarAction(theme.NavigateNextIcon(), func() {})
	t.ZoomOut = widget.NewToolbarAction(theme.ZoomOutIcon(), func() {})
	t.ZoomReset = widget.NewToolbarAction(theme.ZoomFitIcon(), func() {})
	t.ZoomIn = widget.NewToolbarAction(theme.ZoomInIcon(), func() {})

	t.Container = widget.NewToolbar(
		t.Back,
		t.Forward,
		t.ZoomOut,
		t.ZoomReset,
		t.ZoomIn,
	)

	return t
}
