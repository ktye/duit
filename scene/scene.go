// Scene provides a widget for duit, that can render a 3d scene and react to mouse and key events
// with changing the transformation matrix.
//
// The opengl style 3d interface is based on github.com/foglman/{fauxl,ln,pt}.
// For the arcball see: github.com/fogleman/meshview/interactor.go.
package scene

import "github.com/ktye/duitdraw"

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
	Eye       Vector  // Eye position in object space
	Center    Vector  // Centeral position (where the eye looks at) in object space.
	Up        Vector  // Unit direction going up in object space.
	Near, Far float64 // Clipping planes
	Fovy      float64 // Field of view angle in degree in y direction (normally 45..60 deg)

	Width, Height float64 // Destination image dimensions.

	Sensitivity float64
	Start       Vector
	Current     Vector
	Rotation    Matrix
	Translation Vector
	Scroll      float64
	Rotate      bool
	Pan         bool
}

// NewView initializes the view and arcball with default values.
// The values can be changed afterwards.
func NewView() View {
	v := View{
		Eye:  V(-1, -1, 1),
		Up:   V(0, 0, 1),
		Near: 1,
		Far:  100,
		Fovy: 20,

		Sensitivity: 20,
		Rotation:    Identity(),
	}
	return v
}
