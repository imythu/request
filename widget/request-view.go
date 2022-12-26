package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"request/http"
)

func RequestView(r *http.Request) fyne.CanvasObject {

	if r == nil {
		return widget.NewLabel("")
	}
	return widget.NewLabel(r.Url)
}
