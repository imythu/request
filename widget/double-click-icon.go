package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TreeIcon struct {
	widget.Icon
	Uid           string
	OnDoubleClick func(tappable *fyne.PointEvent)
}

func (t *TreeIcon) DoubleTapped(event *fyne.PointEvent) {
	t.OnDoubleClick(event)
}

func NewTappableIcon() *TreeIcon {
	icon := &TreeIcon{}
	icon.ExtendBaseWidget(icon)

	return icon
}
