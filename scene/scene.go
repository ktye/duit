// Scene provides a widget for duit, that can render a 3d scene and react to mouse and key events
// with changing the transformation matrix.
//
// The opengl style 3d interface is based on github.com/foglman/{fauxl,ln,pt}.
// Reactions to mouse events are implemented base on the article by Salvi:
// 	http://salvi.chaosnet.org/texts/gl-view.pdf
// There is also an arcball implementation here: github.com/fogleman/meshview/interactor.go.
package scene

import (
	"fmt"
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
	//Arcball
	Eye           Vector  // Eye position in object space
	Center        Vector  // Centeral position (where the eye looks at) in object space.
	Up            Vector  // Unit direction going up in object space.
	Near, Far     float64 // Clipping planes
	Fovy          float64 // Field of view angle in degree in y direction (normally 45..60 deg)
	Width, Height float64 // Destination image dimensions.
}

func (v View) Matrix() Matrix {
	fmt.Println("ViewMatrix: Eye:", v.Eye, "Center", v.Center)
	return LookAt(v.Eye, v.Center, v.Up).Perspective(v.Fovy, v.Width/v.Height, v.Near, v.Far)
}

func (v *View) Pan(dx, dy float64) {
	length := v.Eye.Sub(v.Center).Length() * 2.0 * math.Tan(v.Fovy*math.Pi/360.0)
	dirx, diry := v.dirxy()
	v.Center = v.Center.Add(diry.MulScalar(dx * length * v.Width / v.Height)).Add(dirx.MulScalar(dy * length))
}

func (v *View) Rotate(dx, dy float64) {
	dirx, diry := v.dirxy()
	v.Eye = rotatePoint(v.Eye, v.Center, dirx, -dx*math.Pi)
	v.Eye = rotatePoint(v.Eye, v.Center, diry, dy*math.Pi)
	v.Up = rotatePoint(v.Center.Add(v.Up), v.Center, diry, dy*math.Pi).Sub(v.Center)
}

// rotatePoint rotates the point around the fixed axis through origin.
func rotatePoint(point, origin, direction Vector, angle float64) Vector {
	// The Rotation matrix may be inverted, compared to the one given by Salvi.
	return origin.Add(Rotate(direction, angle).RotVec(point.Sub(origin)))
}

// RotVec is like MulDirection, but does not normalize the result.
func (a Matrix) RotVec(b Vector) Vector {
	x := a.X00*b.X + a.X01*b.Y + a.X02*b.Z
	y := a.X10*b.X + a.X11*b.Y + a.X12*b.Z
	z := a.X20*b.X + a.X21*b.Y + a.X22*b.Z
	return Vector{x, y, z}
}

// dirxy returns the vectors corresponding to the window axes.
func (v View) dirxy() (Vector, Vector) {
	ce := v.Center.Sub(v.Eye)
	upxce := v.Up.Cross(ce)
	return v.Up, upxce.MulScalar(1.0 / upxce.Length())
}
