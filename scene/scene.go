// Scene provides a widget for duit, that can render a 3d scene and react to mouse and key events
// with changing the transformation matrix.
//
// The opengl style 3d interface is based on github.com/foglman/{fauxl,ln,pt}.
// For the arcball see: github.com/fogleman/meshview/interactor.go.
package scene

import (
	"math"

	"github.com/ktye/duitdraw"
)

// Scene is a UI which renders a 3d scene and reacts to events with pan, zoom and rotate.
type Scene struct {
	View
	SceneDrawer `json:"-"`
}

// SceneDrawer can draw a 3d object onto an image.
type SceneDrawer interface {
	DrawScene(im *duitdraw.Image, view View)
}

// View defines where we look at.
type View struct {
	Eye            Vector  // Eye position in object space
	Center         Vector  // Centeral position (where the eye looks at) in object space.
	Up             Vector  // Unit direction going up in object space.
	Near, Far      float64 // Clipping planes
	Fovy           float64 // Field of view angle in degree in y direction (normally 45..60 deg)
	Width, Height  float64 // Destination image dimensions.
	rotating       bool
	startX, startY float64
}

// NewView initializes the view and arcball with default values.
// The values can be changed afterwards.
func NewView() View {
	v := View{
		Eye:  V(2, 1, 2),
		Up:   V(0, 1, 0),
		Near: 1,
		Far:  100,
		Fovy: 20,
	}
	return v
}

// Zoom moves the camera (Eye) towards or away from the center.
func (v *View) Zoom(scale float64) {
	v.Eye = v.Center.Sub(v.Center.Sub(v.Eye).MulScalar(scale))
}

// Rotate the camera location (eye vector).
// Dx and Dy are relative mouse movements in relation to width and height.
func (v *View) Rotate(dx, dy float64) {
	r := 3 * math.Sqrt(dx*dx+dy*dy)
	x := v.Up.MulScalar(dx / r)
	y := v.Center.Sub(v.Eye).Cross(v.Up).MulScalar(dy / r)
	v.rotateEye(x.Add(y), r)
}

// RotateEye rotates the camera position around the vector a by phi.
// The camera still looks to the center and it's distance remains constant.
func (v *View) rotateEye(a Vector, phi float64) {
	// Don't rotate the camera too far up or down.
	// The scalar product with the Up vector is limited.
	e := v.Center.Sub(Rotate(a, phi).MulPosition(v.Center.Sub(v.Eye)))
	if math.Abs(e.Normalize().Dot(v.Up)) < 0.98 {
		v.Eye = e
	}
}
