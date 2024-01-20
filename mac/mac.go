package mac

import (
	"fmt"
	"net"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var labelMessage *widget.Label

func ShowMacScreen() *fyne.Container {
	labelIP := widget.NewLabel("IP地址")
	txtIP := widget.NewEntry()
	btnGetMac := widget.NewButton("获取MAC地址", func() {
		labelMessage.SetText("MAC地址是：" + "16-F6-D8-2A-DB-EA")
		addr := GetIP()
		txtIP.SetText(addr)
		GetLocalMac()
	})
	labelNull := widget.NewLabel("")
	rightLayout := container.New(layout.NewGridLayout(3), labelIP, txtIP, labelNull, labelNull, btnGetMac)

	labelMessage = widget.NewLabel("扫描结果")

	return container.NewVBox(rightLayout, labelMessage)
}

// 获取本机IP
func GetIP() string {
	dial, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "127.0.0.1"
	}
	addr := dial.LocalAddr().String()
	// fmt.Print(addr)
	index := strings.LastIndex(addr, ":")
	return addr[:index]
}

// 获取本机MAC地址
func GetLocalMac() (mac string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "error"
	}
	for _, inter := range interfaces {
		mac := inter.HardwareAddr
		fmt.Println(mac)
	}
	return mac
}
