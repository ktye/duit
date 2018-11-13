package scene

import (
	"image"
	"math"

	"github.com/ktye/duit"
	"github.com/ktye/duitdraw"
)

// This file implements the duit.UI interface for Scene.

func (s *Scene) Layout(dui *duit.DUI, self *duit.Kid, sizeAvail image.Point, force bool) {
	return
}

func (s *Scene) Draw(dui *duit.DUI, self *duit.Kid, img *duitdraw.Image, orig image.Point, m duitdraw.Mouse, force bool) {
	s.View.Width = float64(img.R.Dx())
	s.View.Height = float64(img.R.Dy())
	s.DrawScene(img, s.View)
}

func (s *Scene) Mouse(dui *duit.DUI, self *duit.Kid, m duitdraw.Mouse, origM duitdraw.Mouse, orig image.Point) (r duit.Result) {
	x, y := float64(m.Point.X), float64(m.Point.Y)
	if m.Buttons == 1 {
		s.View.MouseRotate(x, y, true)
	} else if m.Buttons == 4 {
		s.View.MousePan(x, y, true)
	} else if m.Buttons == 8 {
		s.View.Scroll -= 5
		self.Draw = duit.Dirty
	} else if m.Buttons == 16 {
		s.View.Scroll += 5
		self.Draw = duit.Dirty
	} else {
		// Trigger a redraw after releasing a mouse button.
		if s.View.Rotate || s.View.Pan {
			self.Draw = duit.Dirty
		}
		s.View.MouseRotate(x, y, false)
		s.View.MousePan(x, y, false)
	}

	return r
}

func (s *Scene) Key(dui *duit.DUI, self *duit.Kid, k rune, m duitdraw.Mouse, orig image.Point) (r duit.Result) {
	if k >= '1' && k <= '7' {
		s.View.Translation = Vector{}
		s.View.Scroll = 0
	}
	switch k {
	case '1':
		s.View.Rotation = Identity()
	case '2':
		s.View.Rotation = Identity().Rotate(V(0, 0, 1), math.Pi/2)
	case '3':
		s.View.Rotation = Identity().Rotate(V(0, 0, 1), math.Pi)
	case '4':
		s.View.Rotation = Identity().Rotate(V(0, 0, 1), -math.Pi/2)
	case '5':
		s.View.Rotation = Identity().Rotate(V(1, 0, 0), math.Pi/2)
	case '6':
		s.View.Rotation = Identity().Rotate(V(1, 0, 0), -math.Pi/2)
	case '7':
		s.View.Rotation = Identity().Rotate(V(1, 1, 0).Normalize(), -math.Pi/4).Rotate(V(0, 0, 1), math.Pi/4)
	case duitdraw.KeyLeft:
		s.View.KeyRotate(1, 0)
	case duitdraw.KeyRight:
		s.View.KeyRotate(-1, 0)
	case duitdraw.KeyUp:
		s.View.KeyRotate(0, 1)
	case duitdraw.KeyDown:
		s.View.KeyRotate(0, -1)
	case duitdraw.KeyPageUp:
		s.View.Scroll -= 5
	case duitdraw.KeyPageDown:
		s.View.Scroll += 5
	default:
		return
	}
	self.Draw = duit.Dirty
	return
}

func (s *Scene) FirstFocus(dui *duit.DUI, self *duit.Kid) (warp *image.Point) {
	return nil
}

func (s *Scene) Focus(dui *duit.DUI, self *duit.Kid, o duit.UI) (warp *image.Point) {
	if s != o {
		return nil
	}
	return &image.ZP
}

func (s *Scene) Mark(self *duit.Kid, o duit.UI, forLayout bool) (marked bool) {
	return self.Mark(o, forLayout)
}

func (s *Scene) Print(self *duit.Kid, indent int) {
	duit.PrintUI("Scene", self, indent)
}
