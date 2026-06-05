package views

import (
	"discord-relative-timestamp/consts"
	"discord-relative-timestamp/internal/customtheme"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	slider_Min = 0.5
	slider_Max = 2.5
)

func SetSettingsView(w fyne.Window) {
	btn_exit := widget.NewButton("Save and go Back", nil)
	b1 := container.NewBorder(
		nil, btn_exit,
		nil, nil)

	slider_fyneScale := widget.NewSlider(slider_Min, slider_Max)
	slider_fyneScale.Step = 0.1
	slider_fyneScale.SetValue(fyne.CurrentApp().Preferences().FloatWithFallback(consts.PK_ThemeUI_Scale, 1))

	label_FyneScaleValue := widget.NewLabel(fmt.Sprintf("%.1f", slider_fyneScale.Value))
	slider_fyneScale.OnChanged = func(f float64) {
		label_FyneScaleValue.SetText(fmt.Sprintf("%.1f", f))
	}

	b1.Add(
		container.NewVBox(
			container.NewGridWithColumns(
				2,
				container.NewBorder(
					nil, nil,
					widget.NewLabel("UI Scale"),
					label_FyneScaleValue,
				),
				slider_fyneScale),
			container.NewGridWithColumns(
				2,
				widget.NewButton("LightTheme", func() {
					fyne.CurrentApp().Preferences().SetInt(consts.PK_ThemeVariant, int(theme.VariantLight))
					fyne.CurrentApp().Settings().SetTheme(
						&customtheme.ForcedVariant{
							Theme:   theme.DefaultTheme(),
							Variant: theme.VariantLight,
							Scale:   float32(fyne.CurrentApp().Preferences().Float(consts.PK_ThemeUI_Scale))})
				}),

				widget.NewButton("DarkTheme", func() {
					fyne.CurrentApp().Preferences().SetInt(consts.PK_ThemeVariant, int(theme.VariantDark))
					fyne.CurrentApp().Settings().SetTheme(
						&customtheme.ForcedVariant{
							Theme:   theme.DefaultTheme(),
							Variant: theme.VariantDark,
							Scale:   float32(fyne.CurrentApp().Preferences().Float(consts.PK_ThemeUI_Scale))})
				}),
			),
		),
	)

	btn_exit.OnTapped = func() {
		fyne.CurrentApp().Preferences().SetFloat(consts.PK_ThemeUI_Scale, slider_fyneScale.Value)

		fyne.CurrentApp().Settings().SetTheme(
			&customtheme.ForcedVariant{
				Theme: theme.DefaultTheme(),
				Variant: fyne.ThemeVariant(
					fyne.CurrentApp().Preferences().Float(consts.PK_ThemeVariant)),
				Scale: float32(slider_fyneScale.Value),
			})
		SetMainView(w)
	}
	w.SetContent(b1)
}
