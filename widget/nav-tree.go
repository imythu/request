package widget

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"request/http"
	"request/http/body"
	"request/http/method"
	"request/storage/sqlite3"
)

var data = map[string]interface{}{
	"new-cc": &map[string]interface{}{
		"添加自定义节点": &http.Request{
			Method: method.POST,
			Url:    "http://127.0.0.1:17785/command",
			Params: map[string]string{
				"SoundType": "111",
			},
			Headers: nil,
			Body:    &body.NoneBody{},
		},
	},
}

func Tree() *widget.Tree {
	repo := sqlite3.NewRepo()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			nodes := repo.QueryChildNodes(uid)
			var nodeIdList []string
			for _, node := range nodes {
				nodeIdList = append(nodeIdList, node.ID)
			}
			return nodeIdList
		},
		IsBranch: func(uid string) bool {
			return repo.HasChild(uid)
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t := repo.QueryNode(uid)
			obj.(*widget.Label).SetText(t.Name)
			if repo.HasChild(uid) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Bold: true}
			}
		},
		OnSelected: func(uid string) {
			node := repo.QueryNode(uid)
			fmt.Printf("%v", node)
		},
	}
	tree.Resize(fyne.NewSize(300, 100))
	return tree
}
