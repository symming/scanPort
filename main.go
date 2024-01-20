package main

import (
	"image/color"
	"os"
	"scanPort/host"
	"scanPort/mac"
	"scanPort/scan"
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

var labelStatus *widget.Label

func showMain(w fyne.Window, menu string) {
	//内容
	//w.SetContent(widget.NewLabel("test"))
	labelTop := widget.NewLabel("基于go语言的端口扫描器的设计与实现")

	b1 := widget.NewButton("主机扫描", func() {
		showMain(w, "host")
		labelStatus.SetText("主机扫描...")
	})
	b2 := widget.NewButton("端口扫描", func() {
		showMain(w, "scan")
		labelStatus.SetText("端口扫描...")
	})
	b3 := widget.NewButton("MAC地址获取", func() {
		showMain(w, "mac")
		labelStatus.SetText("获取MAC地址...")

	})
	left := container.NewVBox(b1, b2, b3)

	var rightLayout *fyne.Container
	if menu == "host" {
		rightLayout = host.ShowHostScreen()
	} else if menu == "scan" {
		rightLayout = scan.ShowScanScreen()
	} else if menu == "mac" {
		rightLayout = mac.ShowMacScreen()
	}

	colorGray := color.RGBA{R: 0, G: 0, B: 0, A: 120}
	line1 := canvas.NewLine(colorGray)
	line2 := canvas.NewLine(colorGray)
	labelStatus = widget.NewLabel("主机扫描...")
	right := container.NewVBox(line1, labelTop, line2, labelStatus, rightLayout)

	split := container.NewHSplit(left, right)
	split.Offset = 0.2

	//winmain := container.NewVBox(line1, labelTop, line2, split)
	w.SetContent(split)

}

func main() {
	//新建应用
	a := app.New()
	//新建窗体
	w := a.NewWindow("端口扫描器")
	w.SetMaster()

	showMain(w, "host")

	//设置窗体大小
	w.Resize(fyne.NewSize(900, 600))
	//设置窗口居中
	w.CenterOnScreen()

	//运行程序
	w.ShowAndRun()

}
