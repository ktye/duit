package scene

import "math"

// Arcball changes the scene Matrix reacting to user interactions with mouse and keyboard.
//
// This is adapter from the glfw version:
// github.com/fogleman/meshview/interactor.go
//
// Reference:
// Ken Shoemaker: A User Interface for Specifying Three-Dimensional Orientation Using a Mouse (1992).

func (v *View) KeyRotate(xdir, ydir int) {
	w, h := v.Width, v.Height
	x, y := w/h, h/2
	v.Start = arcballVector(x, y, w, h)
	s := 5.0
	v.Current = arcballVector(x-s*float64(xdir), y-s*float64(ydir), w, h)
	m := arcballRotate(v.Start, v.Current, v.Sensitivity)
	v.Rotation = m.Mul(v.Rotation)
	v.Rotate = false
}

func (v *View) MouseRotate(x, y float64, press bool) {
	w, h := v.Width, v.Height
	if v.Rotate == false && press {
		vec := arcballVector(x, y, w, h)
		v.Start = vec
		v.Current = vec
		v.Rotate = true
	} else if v.Rotate == true && press == false {
		v.Current = arcballVector(x, y, w, h)
		m := arcballRotate(v.Start, v.Current, v.Sensitivity)
		v.Rotation = m.Mul(v.Rotation)
		v.Rotate = false
	}
}

func (v *View) MousePan(x, y float64, press bool) {
	w, h := v.Width, v.Height
	if v.Pan == false && press {
		vec := screenPosition(x, y, w, h)
		v.Start = vec
		v.Current = vec
		v.Pan = true
	} else if v.Pan == true && press == false {
		v.Current = screenPosition(x, y, w, h)
		d := v.Current.Sub(v.Start)
		v.Translation = v.Translation.Add(d)
		v.Pan = false
	}
}

func (v *View) Matrix() Matrix {
	r := v.Rotation
	if v.Rotate {
		r = arcballRotate(v.Start, v.Current, v.Sensitivity).Mul(r)
	}
	t := v.Translation
	if v.Pan {
		t = t.Add(v.Current.Sub(v.Start))
	}
	s := math.Pow(0.98, v.Scroll)
	m := Identity()
	m = m.Scale(V(s, s, s))
	m = r.Mul(m)
	m = m.Translate(t)
	m = m.LookAt(v.Eye, v.Center, v.Up)
	m = m.Perspective(v.Fovy, v.Width/v.Height, v.Near, v.Far)
	return m
}

func screenPosition(x, y, w, h float64) Vector {
	x = (x/w)*2 - 1
	y = (y/h)*2 - 1
	return Vector{x, 0, -y}
}

func arcballVector(x, y, w, h float64) Vector {
	x = (x/w)*2 - 1
	y = (y/h)*2 - 1
	x /= 4
	y /= 4
	x = -x
	q := x*x + y*y
	if q <= 1 {
		z := math.Sqrt(1 - q)
		return Vector{x, z, y}
	} else {
		return Vector{x, 0, y}.Normalize()
	}
}

func arcballRotate(a, b Vector, sensitivity float64) Matrix {
	const eps = 1e-9
	dot := b.Dot(a)
	if math.Abs(dot) < eps || math.Abs(dot-1) < eps {
		return Identity()
	} else if math.Abs(dot+1) < eps {
		return Rotate(a.Perpendicular(), math.Pi*sensitivity)
	} else {
		angle := math.Acos(dot)
		v := b.Cross(a).Normalize()
		return Rotate(v, angle*sensitivity)
	}
}
