package ui

import (
	"bytes"
	"image"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
)

var uidata struct {
	qr_box     *fyne.Container
	ip_lable   *widget.Label
	selectd_ip string
	ip         []string
}

func Start(ip []string) {
	uidata.ip = ip
	uidata.qr_box = NewSizedBox(300, 300, false)
	uidata.ip_lable = widget.NewLabel("")

	for _, ip_ := range ip {
		if strings.HasPrefix(ip_, "http://192") {
			uidata.selectd_ip = ip_
		}
	}

	a := app.New()
	w := a.NewWindow("文件传输助手")
	w.Resize(fyne.Size{Width: 600, Height: 400})
	w.SetFixedSize(true)
	vbox := container.New(layout.NewVBoxLayout())
	hbox := container.New(layout.NewHBoxLayout())

	ip_list := widget.NewList(
		func() int {
			return len(ip)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(ip[lii])

		})
	ip_list.OnSelected = func(id widget.ListItemID) {
		uidata.selectd_ip = uidata.ip[id]
		new_qr_image()
	}

	ip_box := NewSizedBox(300, 350, false)
	ip_box.Add(ip_list)
	vbox.Add(widget.NewLabel("Welcome"))
	vbox.Add(hbox)
	hbox.Add(ip_box)
	vvbox := container.New(layout.NewVBoxLayout())
	vvbox.Add(uidata.qr_box)
	hhbox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), uidata.ip_lable, layout.NewSpacer())

	vvbox.Add(hhbox)
	hbox.Add(vvbox)
	// b := container.New(layout.NewVBoxLayout(), widget.NewLabel("hello"), image)

	w.SetContent(vbox)

	if len(uidata.selectd_ip) > 0 {
		new_qr_image()
	}

	w.ShowAndRun()
}

func new_qr_image() {
	i, e := createQR(uidata.selectd_ip)
	if e != nil {
		return
	}
	uidata.qr_box.RemoveAll()
	uidata.qr_box.Add(canvas.NewImageFromImage(i))
	uidata.ip_lable.SetText(uidata.selectd_ip)
}

func createQR(s string) (image.Image, error) {
	b, e := qrcode.Encode(s, qrcode.Medium, 256)
	if e != nil {
		return nil, e
	}
	img, _, e := image.Decode(bytes.NewReader(b))
	if e != nil {
		return nil, e
	}
	return img, nil
}
