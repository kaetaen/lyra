package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kaetaen/lyra/theming"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theming.LyraTheme{})

	w := a.NewWindow("Lyra")
	w.Resize(fyne.NewSize(620, 620))

	lyricText := widget.NewLabel("Letra de Música")
	lyricText.Wrapping = fyne.TextWrapWord
	lyricText.Alignment = fyne.TextAlignCenter

	content := container.NewVScroll(lyricText)
	content.SetMinSize(fyne.NewSize(600, 600))

	screen := container.New(layout.NewCenterLayout(), content)

	w.SetContent(screen)

	w.ShowAndRun()
}
