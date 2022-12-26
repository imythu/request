package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type TreeIcon struct {
	widget.Icon
	Uid           string
	OnDoubleClick func(tappable *fyne.PointEvent)
}

func (t *TreeIcon) TappedSecondary(event *fyne.PointEvent) {
	menu := fyne.NewMenu("右键",
		fyne.NewMenuItem("删除", func() {
			dialog.ShowConfirm("警告", fmt.Sprintf("确认删除 uid %s 吗？", t.Uid), func(b bool) {
				if b {
					repo.DeleteNode(t.Uid)
				}
			}, fyne.CurrentApp().Driver().AllWindows()[0])
		}))
	entryPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(t)
	popUpPos := entryPos.Add(event.Position)
	c := fyne.CurrentApp().Driver().CanvasForObject(t)
	widget.ShowPopUpMenuAtPosition(menu, c, popUpPos)
}

func (t *TreeIcon) DoubleTapped(event *fyne.PointEvent) {
	t.OnDoubleClick(event)
}

func NewTappableIcon() *TreeIcon {
	icon := &TreeIcon{}
	icon.ExtendBaseWidget(icon)

	return icon
}
