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
	label := widget.NewLabel("pdfview - no file loaded")
	w.SetContent(container.NewCenter(label))
	w.ShowAndRun()
}
