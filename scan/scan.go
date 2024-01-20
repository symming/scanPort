package scan

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var labelMessage *widget.Label

func ShowScanScreen() *fyne.Container {
	labelStartIP := widget.NewLabel("起始IP地址")
	txtStartIP := widget.NewEntry()
	labelStartPort := widget.NewLabel("开始端口")
	txtStartPort := widget.NewEntry()
	labelEndIP := widget.NewLabel("结束IP地址")
	txtEndIP := widget.NewEntry()
	labelEndPort := widget.NewLabel("结束端口")
	txtEndPort := widget.NewEntry()

	btnStart := widget.NewButtonWithIcon("开始扫描", theme.MediaPlayIcon(), func() {
		labelMessage.SetText("正在扫描端口...")
	})
	btnEnd := widget.NewButtonWithIcon("停止扫描", theme.MediaPauseIcon(), func() {
		labelMessage.SetText("端口扫描结果是...")
	})
	labelNull := widget.NewLabel("")

	rightLayout := container.New(layout.NewGridLayout(5), labelStartIP, txtStartIP, labelStartPort, txtStartPort,
		labelNull, labelEndIP, txtEndIP, labelEndPort, txtEndPort, labelNull, labelNull, btnStart, labelNull, btnEnd)

	labelMessage = widget.NewLabel("扫描结果")

	return container.NewVBox(rightLayout, labelMessage)
}
