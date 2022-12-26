package resource

import (
	"bufio"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"io"
	"os"
)

type MonoFontTheme struct{}

var monoLight = getResource("./ttf/LXGWWenKaiMono-Light.ttf")
var monoBold = getResource("./ttf/LXGWWenKaiMono-Bold.ttf")

func (t *MonoFontTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (t *MonoFontTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t *MonoFontTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (*MonoFontTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	// 此处可以根据不同字重返回不同的字体, 但是我用的都是同样的字体
	if s.Bold {
		if s.Italic {
			return monoBold
		}
		// 返回自定义字体
		return monoBold
	}
	if s.Italic {
		return monoBold
	}
	// 返回自定义字体
	return monoLight
}

func getResource(resourcePath string) *fyne.StaticResource {
	finalByte := make([]byte, 0)

	fi, err := os.Open(resourcePath)

	if err != nil {
		panic(err)
	}

	defer fi.Close()
	r := bufio.NewReader(fi)

	readBuf := make([]byte, 1024)
	for {
		n, err := r.Read(readBuf)
		if err != nil && err != io.EOF {
			panic(err)
			//return
		}
		if 0 == n {
			break
		} else {
			finalByte = append(finalByte, readBuf[:n]...)
		}
	}
	fmt.Printf("load resource :" + resourcePath + " success !")

	return &fyne.StaticResource{
		StaticName:    fi.Name(),
		StaticContent: finalByte,
	}
}
