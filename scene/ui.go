package scene

import (
	"fmt"
	"image"

	"github.com/ktye/duit"
	"github.com/ktye/duitdraw"
)

// This file implements the duit.UI interface for Scene.

func (s *Scene) Layout(dui *duit.DUI, self *duit.Kid, sizeAvail image.Point, force bool) {
	fmt.Println("Layout sizeAvail", sizeAvail)
	return
}

func (s *Scene) Draw(dui *duit.DUI, self *duit.Kid, img *duitdraw.Image, orig image.Point, m duitdraw.Mouse, force bool) {
	fmt.Println("Draw")
	s.View.Width = float64(img.R.Dx())
	s.View.Height = float64(img.R.Dy())
	s.DrawScene(img, s.View)
}

func (s *Scene) Mouse(dui *duit.DUI, self *duit.Kid, m duitdraw.Mouse, origM duitdraw.Mouse, orig image.Point) (r duit.Result) {
	fmt.Println("Mouse", m.Buttons)

	// Mouse wheel zooms in and out.
	// It moves the Eye position towards the Center.
	// The amount is set to 5% of the distance.
	if m.Buttons == 8 || m.Buttons == 16 {
		d := s.View.Center.Sub(s.View.Eye).MulScalar(0.05)
		if m.Buttons == 16 {
			d = d.Negate()
		}
		s.View.Eye = s.View.Eye.Add(d)
		r.Consumed = true
		self.Draw = duit.Dirty
	}

	// TODO: mouse events for zoom and pan.
	// TODO: we cannot use the shift key or other modifiers in combination with mouse movements.
	// It is not present in duit.
	// We have to use another key and detect key press individually from mouse events.

	return r
}

func (s *Scene) Key(dui *duit.DUI, self *duit.Kid, k rune, m duitdraw.Mouse, orig image.Point) (r duit.Result) {
	fmt.Println("Key", k, duitdraw.KeyLeft, duitdraw.KeyRight, duitdraw.KeyDown, duitdraw.KeyUp)

	switch k {
	case duitdraw.KeyLeft:
		s.Pan(-0.1, 0)
	case duitdraw.KeyRight:
		s.Pan(0.1, 0)
	case duitdraw.KeyUp:
		s.Pan(0, -0.1)
	case duitdraw.KeyDown:
		s.Pan(0, 0.1)
	case 'j':
		s.Rotate(0.01, 0)
	case 'l':
		s.Rotate(-0.01, 0)
	case 'i':
		s.Rotate(0, 0.02)
	case 'm':
		s.Rotate(0, -0.02)
	default:
		fmt.Println(string([]rune{k}))
		return
	}
	self.Draw = duit.Dirty
	return
}

func (s *Scene) FirstFocus(dui *duit.DUI, self *duit.Kid) (warp *image.Point) {
	fmt.Println("FirstFocus")
	return nil
}

func (s *Scene) Focus(dui *duit.DUI, self *duit.Kid, o duit.UI) (warp *image.Point) {
	fmt.Println("Focus")
	if s != o {
		return nil
	}
	return &image.ZP
}

func (s *Scene) Mark(self *duit.Kid, o duit.UI, forLayout bool) (marked bool) {
	fmt.Println("Mark")
	return self.Mark(o, forLayout)
}

func (s *Scene) Print(self *duit.Kid, indent int) {
	fmt.Println("Print")
	duit.PrintUI("Scene", self, indent)
}
