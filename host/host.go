package host

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var labelMessage *widget.Label

func ShowHostScreen() *fyne.Container {
	labelStartIP := widget.NewLabel("起始IP地址")
	txtStartIP := widget.NewEntry()
	labelEndIP := widget.NewLabel("结束IP地址")
	txtEndIP := widget.NewEntry()
	btnStart := widget.NewButtonWithIcon("开始扫描", theme.MediaPlayIcon(), func() {
		labelMessage.SetText("正在扫描...")
	})
	btnEnd := widget.NewButtonWithIcon("停止扫描", theme.MediaPauseIcon(), func() {
		labelMessage.SetText("扫描结果是...")
	})
	labelNull := widget.NewLabel("")

	rightLayout := container.New(layout.NewGridLayout(3), labelStartIP, txtStartIP, labelNull,
		labelEndIP, txtEndIP, labelNull, btnStart, btnEnd)

	labelMessage = widget.NewLabel("扫描结果")

	return container.NewVBox(rightLayout, labelMessage)
}
