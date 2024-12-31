package theming

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type LyraTheme struct{}

func (m LyraTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m LyraTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 24
	}
	return theme.DefaultTheme().Size(name)
}

func (m LyraTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m LyraTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
