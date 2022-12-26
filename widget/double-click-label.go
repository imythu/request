package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TreeLabel struct {
	widget.Label
	Uid           string
	OnDoubleClick func(tappable *fyne.PointEvent)
}

func (t *TreeLabel) DoubleTapped(event *fyne.PointEvent) {
	t.OnDoubleClick(event)
}

func newTreeLabel(text string) *TreeLabel {
	label := &TreeLabel{}
	label.ExtendBaseWidget(label)
	label.SetText(text)

	return label
}
