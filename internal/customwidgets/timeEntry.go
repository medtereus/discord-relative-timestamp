package customwidgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewCustomEntry(label string) *fyne.Container {
	e := newNumericalEntry()
	e.MultiLine = false
	e.Scroll = container.ScrollNone
	l := widget.NewLabel(label)

	l.SizeName = theme.SizeNameInnerPadding

	return container.New(&timeEntryLayout{}, e, l)
}

type timeEntryLayout struct{}

func (l *timeEntryLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	entry := objects[0]
	label := objects[1]

	entry.Resize(size)

	ls := label.MinSize()
	label.Move(fyne.NewPos(
		size.Width-ls.Width,
		size.Height-ls.Height,
	))
	label.Resize(ls)
}

func (l *timeEntryLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return objects[0].MinSize()
}
