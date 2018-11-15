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
	if m.Buttons == 8 {
		s.View.Zoom(1 / 1.1)
		self.Draw = duit.Dirty
	} else if m.Buttons == 16 {
		s.View.Zoom(1.1)
		self.Draw = duit.Dirty
	} else if m.Buttons == 1 {
		if s.rotating == false {
			s.startX = x
			s.startY = y
			s.rotating = true
		}
	} else if m.Buttons == 0 {
		if s.rotating {
			s.rotating = false
			dx := (x - s.startX) / s.View.Width
			dy := (y - s.startY) / s.View.Height
			s.View.Rotate(dx, dy)
			self.Draw = duit.Dirty
		}
	}
	return r
}

func (s *Scene) Key(dui *duit.DUI, self *duit.Kid, k rune, m duitdraw.Mouse, orig image.Point) (r duit.Result) {
	switch k {
	case '1', '2', '3', '4', '5', '6', '7', '8':
		x := k - '1'
		e1, e2 := s.Up.Perpendicular(), s.Up.Perpendicular().Cross(s.Up)
		views := []Vector{e1, e2, e1.Negate(), e2.Negate(), e1.Add(e2), e2.Sub(e1), e1.Add(e2).Negate(), e1.Sub(e2)}
		dist := s.Center.Sub(s.Eye).Length()
		if x < 4 {
			s.Eye = views[x].MulScalar(dist)
		} else {
			s.Eye = views[x].Add(s.Up).MulScalar(dist / math.Sqrt(3))
		}
	case 'x':
		s.Up = V(1, 0, 0)
	case 'y':
		s.Up = V(0, 1, 0)
	case 'z':
		s.Up = V(0, 0, 1)
	case duitdraw.KeyLeft:
		s.Rotate(-0.05, 0)
	case duitdraw.KeyRight:
		s.View.Rotate(0.05, 0)
	case duitdraw.KeyUp:
		s.View.Rotate(0, -0.05)
	case duitdraw.KeyDown:
		s.View.Rotate(0, 0.05)
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
}
