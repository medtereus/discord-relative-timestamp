package customtheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type ForcedVariant struct {
	fyne.Theme

	Variant fyne.ThemeVariant
	Scale   float32
}

func (f *ForcedVariant) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return f.Theme.Color(name, f.Variant)
}
func (f *ForcedVariant) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name) * f.Scale
}
