package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"request/storage/sqlite3"
)

const uidKey = "uid"

func NavTree() fyne.CanvasObject {
	repo := sqlite3.NewRepo()
	tree := &widget.Tree{
		Root: "",
	}
	tree.IsBranch = func(uid string) bool {
		return repo.HasChild(uid)
	}
	tree.ChildUIDs = func(uid string) []string {
		nodes := repo.QueryChildNodes(uid)
		var nodeIdList []string
		for _, node := range nodes {
			nodeIdList = append(nodeIdList, node.ID)
		}
		return nodeIdList
	}
	tree.CreateNode = func(branch bool) fyne.CanvasObject {
		var icon fyne.CanvasObject
		if branch {
			icon = widget.NewIcon(nil)
		} else {
			icon = widget.NewFileIcon(nil)
		}
		label := widget.NewLabel("Loading")

		border := container.NewBorder(nil, nil, icon, nil, label)
		return border
	}
	tree.UpdateNode = func(uid string, branch bool, obj fyne.CanvasObject) {
		node := repo.QueryNode(uid)
		wrapper := obj.(*doubleTappableWidget)
		container := wrapper.Container
		label := container.Objects[0].(*widget.Label)
		label.SetText(node.Name)
		wrapper.data[uidKey] = uid
		if branch {
			label.TextStyle = fyne.TextStyle{Bold: true}
		}
		if branch {
			var r fyne.Resource
			if tree.IsBranchOpen(uid) {
				// Set open folder icon
				r = theme.FolderOpenIcon()
			} else {
				// Set folder icon
				r = theme.FolderIcon()
			}
			container.Objects[1].(*widget.Icon).SetResource(r)
		}
	}

	return tree
}

func (d *doubleTappableWidget) DoubleTapped(event *fyne.PointEvent) {
	d.doubleTapped(event)
}

type doubleTappableWidget struct {
	widget.BaseWidget
	data         map[string]string
	doubleTapped func(*fyne.PointEvent)
}

func (d *doubleTappableWidget) CreateRenderer() fyne.WidgetRenderer {
	d.BaseWidget.ExtendBaseWidget(d)
	return d.BaseWidget
}

type doubleTappableIcon struct {
	doubleTappableWidget
}

func newDoubleTappableIcon() *widget.Icon {
	icon := &doubleTappableIcon{}
	icon.ExtendBaseWidget(icon)
	return icon
}
