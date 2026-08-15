package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type terminalTheme struct{}

var (
	bgColor  = color.NRGBA{R: 0x0d, G: 0x0d, B: 0x0d, A: 0xff}
	fgColor  = color.NRGBA{R: 0x33, G: 0xff, B: 0x66, A: 0xff}
	dimColor = color.NRGBA{R: 0x1a, G: 0x1a, B: 0x1a, A: 0xff}
)

func (terminalTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return bgColor
	case theme.ColorNameForeground:
		return fgColor
	case theme.ColorNameButton, theme.ColorNameDisabledButton:
		return dimColor
	default:
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

func (terminalTheme) Font(style fyne.TextStyle) fyne.Resource {
	style.Monospace = true
	return theme.DefaultTheme().Font(style)
}

func (terminalTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (terminalTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func newDefaultSize() fyne.Size {
	return fyne.NewSize(900, 700)
}
