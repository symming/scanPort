package main

import (
	"image/color"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

func init() {
	//设置中文字体:解决中文乱码问题
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "msyh.ttf") || strings.Contains(path, "simhei.ttf") || strings.Contains(path, "simsun.ttc") || strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

var r1 *widget.Label

func main() {
	//新建应用
	a := app.New()
	//新建窗体
	w := a.NewWindow("端口扫描器")
	w.SetMaster()

	//内容
	//w.SetContent(widget.NewLabel("test"))
	labelTop := widget.NewLabel("基于go语言的端口扫描器的设计与实现")

	b1 := widget.NewButton("主机扫描", func() {
		r1.SetText("gaibianle ")
	})
	b2 := widget.NewButton("端口扫描", func() {
		r1.SetText(".....")
	})
	b3 := widget.NewButton("MAC地址获取", func() {

	})
	left := container.NewVBox(b1, b2, b3)

	r1 = widget.NewLabel("扫描")
	r2 := widget.NewLabel("主机")
	right := container.NewVBox(r1, r2)

	split := container.NewHSplit(left, right)
	split.Offset = 0.2

	colorGray := color.RGBA{R: 0, G: 0, B: 0, A: 120}
	line1 := canvas.NewLine(colorGray)
	line2 := canvas.NewLine(colorGray)

	winmain := container.NewVBox(line1, labelTop, line2, split)
	w.SetContent(winmain)

	//设置窗体大小
	w.Resize(fyne.NewSize(900, 600))
	//设置窗口居中
	w.CenterOnScreen()

	//运行程序
	w.ShowAndRun()

}
