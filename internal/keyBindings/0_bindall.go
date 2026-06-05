package keybindings

import (
	"fyne.io/fyne/v2"
)

var all_binds = make([]func(*fyne.KeyEvent), 0)

func ApplyBinds_MainWindow(w fyne.Window) {
	for _, bindingF := range all_binds {
		w.Canvas().SetOnTypedKey(bindingF)
	}
}
