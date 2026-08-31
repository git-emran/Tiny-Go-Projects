package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.NewWithID("com.example.pdfviewer")
	a.Settings().SetTheme(&terminalTheme{})

	w := a.NewWindow("pdfview")
	w.Resize(newDefaultSize())

	engine, err := newEngine()
	if err != nil {
		log.Fatal(err)
	}
	defer engine.Close()

	pageCount, err := engine.OpenFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	viewer := newViewer(engine, pageCount)
	tb := newToolbar()
	tb.Back.OnActivated = func() { logErr(viewer.Prev()) }
	tb.Forward.OnActivated = func() { logErr(viewer.Next()) }
	tb.ZoomIn.OnActivated = func() { logErr(viewer.ZoomIn()) }
	tb.ZoomOut.OnActivated = func() { logErr(viewer.ZoomOut()) }
	tb.ZoomReset.OnActivated = func() { logErr(viewer.ZoomReset()) }

	content := container.NewBorder(tb.Container, nil, nil, nil, viewer.image)
	w.SetContent(content)

	w.Canvas().SetOnTypedKey(func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyRight, "L":
			logErr(viewer.Next())

		case fyne.KeyLeft, "H":
			logErr(viewer.Prev())

		case fyne.KeyPlus, "Equal":
			logErr(viewer.ZoomIn())

		case fyne.KeyMinus:
			logErr(viewer.ZoomIn())
		}
	})

	if err := viewer.render(); err != nil {
		log.Fatal(err)
	}

	w.ShowAndRun()
}

func logErr(err error) {
	if err != nil {
		log.Println("error:", err)
	}
}
