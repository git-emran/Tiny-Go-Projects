package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("com.example.pdfviewer")
	a.Settings().SetTheme(&terminalTheme{})

	w := a.NewWindow("pdfview")
	w.Resize(newDefaultSize())

	tb := newToolbar()
	viewer := newViewer()
	content := container.NewBorder(tb.Container, nil, nil, nil, viewer.image)
	w.SetContent(content)

	if len(os.Args) > 1 {
		engine, err := newEngine()
		if err != nil {
			log.Fatal(err)
		}

		defer engine.Close()

		if _, err := engine.OpenFile(os.Args[1]); err != nil {
			log.Fatal(err)
		}
		page, err := engine.RenderPage(0, 1.0)
		if err != nil {
			log.Fatal(err)
		}
		viewer.Show(page)
	}

	w.ShowAndRun()
}
