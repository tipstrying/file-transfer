package ui

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type SizedBox struct {
	width        float32
	height       float32
	minSizeChild bool
}

func (d *SizedBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(d.width, d.height)
}
func (d *SizedBox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) > 0 {
		if d.minSizeChild {
			//	h_offset := (containerSize.Height - objects[0].MinSize().Height) / 2
			//	w_offset := (containerSize.Width - objects[0].MinSize().Width) / 2
			pos := fyne.NewPos(0, 0)
			//objects[0].Resize(fyne.Size{Width: d.width, Height: d.height})
			// objects[0].MinSize()
			objects[0].Move(pos)
			return
		}
		objects[0].Resize(fyne.Size{Width: d.width, Height: d.height})
		objects[0].Move(fyne.NewPos(0, 0))
		// println(data.windows.Canvas().Content().Size().Width)
		// pos = pos.Add(fyne.NewPos(size.Width, size.Height))
	}
}

func (d *SizedBox) Tapped(_ *fyne.PointEvent) {
	log.Println("I have been tapped")
}

func NewSizedBox(w float32, h float32, minsize bool) *fyne.Container {
	return container.New(&SizedBox{width: w, height: h, minSizeChild: minsize})
}
func NewSizedBoxWithChild(w float32, h float32, minsize bool, child fyne.CanvasObject) *fyne.Container {

	c := container.New(&SizedBox{width: w, height: h, minSizeChild: minsize})
	c.Add(child)
	return c
}

type onTapScroll func()
type TapScroll struct {
	container.Scroll
	onTap onTapScroll
}

func newTappableIcon(content fyne.CanvasObject, tap onTapScroll) *TapScroll {
	icon := &TapScroll{}
	icon.Direction = container.ScrollBoth
	icon.Content = content
	icon.onTap = tap
	icon.ExtendBaseWidget(icon)

	return icon
}
func (t *TapScroll) Tapped(_ *fyne.PointEvent) {
	println("tapped")
	t.onTap()
}
