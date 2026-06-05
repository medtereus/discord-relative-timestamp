package keybindings

import (
	"time"

	"fyne.io/fyne/v2"
)

func init() {
	all_binds = append(all_binds, kb_quit_onDoubleESC)
}

var escTimestamp time.Time

func kb_quit_onDoubleESC(ke *fyne.KeyEvent) {
	if ke.Name == fyne.KeyEscape {
		if time.Since(escTimestamp) < time.Second/2 {
			fyne.CurrentApp().Quit()
		}
		escTimestamp = time.Now()
	}
}
