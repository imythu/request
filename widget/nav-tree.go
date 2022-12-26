package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"request/storage/sqlite3"
)

var repo *sqlite3.NodeRepo

func NavTree() fyne.CanvasObject {
	repo = sqlite3.NewRepo()
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
		icon := NewTappableIcon()
		label := newTreeLabel("Loading")
		f := func(event *fyne.PointEvent) {
			if branch {
				uid := icon.Uid
				if tree.IsBranchOpen(uid) {
					tree.CloseBranch(uid)
				} else {
					tree.OpenBranch(uid)
				}
			}
		}

		icon.OnDoubleClick = f
		label.OnDoubleClick = f

		border := container.NewBorder(nil, nil, icon, nil, label)
		return border
	}
	tree.UpdateNode = func(uid string, branch bool, obj fyne.CanvasObject) {
		node := repo.QueryNode(uid)
		c := obj.(*fyne.Container)
		label := c.Objects[0].(*TreeLabel)
		label.SetText(node.Name)
		label.Uid = uid
		if branch {
			label.TextStyle = fyne.TextStyle{Bold: true}
		}
		icon := c.Objects[1].(*TreeIcon)
		if branch {
			var r fyne.Resource
			if tree.IsBranchOpen(uid) {
				// Set open folder icon
				r = theme.FolderOpenIcon()
			} else {
				// Set folder icon
				r = theme.FolderIcon()
			}
			icon.SetResource(r)
			icon.Uid = uid
		} else {
			icon.SetResource(theme.FileIcon())
		}
	}

	return tree
}
