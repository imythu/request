package main

import (
	"fyne.io/fyne/v2"
	"io/ioutil"
	"request/resource"
	"request/widget"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&resource.MonoFontTheme{})
	myWindow := myApp.NewWindow("Border Layout")
	//bgImg := &canvas.Image{
	//	Resource:     fyne.NewStaticResource("bg", fileToBytes("./2560-1440.png")),
	//	Translucency: 0.45}
	//content := container.NewMax(bgImg)
	//content.Add(widget.Tree())

	content := container.NewMax(widget.Tree())
	mainContainer := container.New(layout.NewBorderLayout(nil, nil, nil, nil), content)
	myWindow.SetContent(mainContainer)
	myWindow.Resize(fyne.NewSize(2560/2, 1440/2))
	myWindow.ShowAndRun()
}

func fileToBytes(filepath string) []byte {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return bytes
}
