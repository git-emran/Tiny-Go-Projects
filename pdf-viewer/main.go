package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	a.Settings().SetTheme(&terminalTheme{})

	w := a.NewWindow("pdfview")
	w.Resize(newDefaultSize())
	tb := newToolbar()
	label := widget.NewLabel("pdfview - no file loaded")
	content := container.NewBorder(tb.Container, nil, nil, nil, container.NewCenter(label))
	w.SetContent(content)
	w.ShowAndRun()
}
