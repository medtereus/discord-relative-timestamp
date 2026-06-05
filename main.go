package main

import (
	"discord-relative-timestamp/consts"
	"discord-relative-timestamp/internal/customtheme"
	keybindings "discord-relative-timestamp/internal/keyBindings"
	"discord-relative-timestamp/internal/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	AppInstance := app.NewWithID(consts.APP_ID)
	AppInstance.Settings().SetTheme(
		&customtheme.ForcedVariant{
			Theme: theme.DefaultTheme(),
			Variant: fyne.ThemeVariant(
				AppInstance.Preferences().Int(consts.PK_ThemeVariant)),
			Scale: float32(AppInstance.Preferences().FloatWithFallback(consts.PK_ThemeUI_Scale, 1))})

	mainWindow := AppInstance.NewWindow(consts.APP_NAME)

	keybindings.ApplyBinds_MainWindow(mainWindow)

	views.SetMainView(mainWindow)

	mainWindow.ShowAndRun()
}
