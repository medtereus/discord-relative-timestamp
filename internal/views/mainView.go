package views

import (
	"discrodRelativeTimestamp/consts"
	"discrodRelativeTimestamp/internal/customwidgets"
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const cb_ON = "Mode: Add to currentTime"
const cb_OFF = "Mode: Sub from currentTime"

func SetMainView(w fyne.Window) {
	cb_CheckTime := widget.NewCheck(
		cb_ON,
		nil)
	cb_CheckTime.SetChecked(
		fyne.CurrentApp().Preferences().BoolWithFallback(consts.PK_TimeCB, true))

	//

	btn_Settings := widget.NewButtonWithIcon(
		"",
		theme.SettingsIcon(),
		func() { SetSettingsView(w) })

	entry_output := widget.NewEntry()
	bottom := container.NewVBox(
		container.NewBorder(
			nil, nil, nil,
			widget.NewButtonWithIcon(
				"Copy",
				theme.ContentCopyIcon(),
				func() { fyne.CurrentApp().Clipboard().SetContent(entry_output.Text) }),
			entry_output))

	border_root := container.NewBorder(
		container.NewBorder(nil, nil, cb_CheckTime, btn_Settings),
		bottom, nil, nil)

	//
	var (
		e_Days    = customwidgets.NewCustomEntry("Days")
		e_Hours   = customwidgets.NewCustomEntry("Hours")
		e_Minutes = customwidgets.NewCustomEntry("Minutes")
		e_Seconds = customwidgets.NewCustomEntry("Seconds")
	)

	collectData := func() time.Duration {
		d, err := strconv.Atoi(e_Days.Objects[0].(*customwidgets.NumericalEntry).Text)
		if err != nil {
			d = 0
		}

		h, err := strconv.Atoi(e_Hours.Objects[0].(*customwidgets.NumericalEntry).Text)
		if err != nil {
			h = 0
		}
		h = h + (d * 24)

		m, err := strconv.Atoi(e_Minutes.Objects[0].(*customwidgets.NumericalEntry).Text)
		if err != nil {
			m = 0
		}
		s, err := strconv.Atoi(e_Seconds.Objects[0].(*customwidgets.NumericalEntry).Text)
		if err != nil {
			s = 0
		}

		Dur, err := time.ParseDuration(
			fmt.Sprintf("%dh%dm%ds", h, m, s))
		if err != nil {
			println("Error in parseDuration:" + err.Error())
			return time.Hour
		}

		return Dur
	}

	updateOutputEntry := func(s string) {
		enteredTime := collectData()

		var T time.Time
		if fyne.CurrentApp().Preferences().Bool(consts.PK_TimeCB) {
			T = time.Now().Add(enteredTime)
		} else {
			T = time.Now().Add(enteredTime * -1)
		}

		entry_output.SetText(fmt.Sprintf("<t:%d:R>", T.Unix()))
	}

	e_Days.Objects[0].(*customwidgets.NumericalEntry).OnChanged = updateOutputEntry
	e_Hours.Objects[0].(*customwidgets.NumericalEntry).OnChanged = updateOutputEntry
	e_Minutes.Objects[0].(*customwidgets.NumericalEntry).OnChanged = updateOutputEntry
	e_Seconds.Objects[0].(*customwidgets.NumericalEntry).OnChanged = updateOutputEntry

	cb_CheckTime.OnChanged = func(b bool) {
		fyne.CurrentApp().Preferences().SetBool(consts.PK_TimeCB, b)

		updateOutputEntry("")

		switch b {
		case true:
			cb_CheckTime.SetText(cb_ON)
		case false:
			cb_CheckTime.SetText(cb_OFF)
		}
	}
	//

	border_root.Add(
		container.NewBorder(
			container.NewGridWithColumns(
				4,
				e_Days, e_Hours, e_Minutes, e_Seconds),
			nil, nil, nil))

	w.SetContent(border_root)
}
