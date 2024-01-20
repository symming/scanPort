package scan

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var labelMessage *widget.Label

// 显示端口扫描布局
func ShowScanScreen() *fyne.Container {
	labelStartIP := widget.NewLabel("起始IP地址")
	txtStartIP := widget.NewEntry()
	labelStartPort := widget.NewLabel("开始端口")
	txtStartPort := widget.NewEntry()
	labelEndIP := widget.NewLabel("结束IP地址")
	txtEndIP := widget.NewEntry()
	labelEndPort := widget.NewLabel("结束端口")
	txtEndPort := widget.NewEntry()

	//开始扫描按钮
	btnStart := widget.NewButtonWithIcon("开始扫描", theme.MediaPlayIcon(), func() {
		labelMessage.SetText("正在扫描端口...")
	})
	//结束扫描按钮
	btnEnd := widget.NewButtonWithIcon("停止扫描", theme.MediaPauseIcon(), func() {
		labelMessage.SetText("端口扫描结果是...")
	})
	labelNull := widget.NewLabel("")

	//布局
	rightLayout := container.New(layout.NewGridLayout(5), labelStartIP, txtStartIP, labelStartPort, txtStartPort,
		labelNull, labelEndIP, txtEndIP, labelEndPort, txtEndPort, labelNull, labelNull, btnStart, labelNull, btnEnd)

	//扫描结果
	labelMessage = widget.NewLabel("扫描结果")

	return container.NewVBox(rightLayout, labelMessage)
}
