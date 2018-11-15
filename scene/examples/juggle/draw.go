package main

import (
	"image"
	imdraw "image/draw"

	"github.com/fogleman/fauxgl"
	"github.com/ktye/duit/scene"
	"github.com/ktye/duitdraw"
)

// Draw implements a scene.SceneDrawer.
// It has the same signature as the redraw function, which is converted to a SceneDrawer by type casting.
type draw func(im *duitdraw.Image, matrix fauxgl.Matrix, eye fauxgl.Vector)

func (d draw) DrawScene(im *duitdraw.Image, v scene.View) {
	m := scene.LookAt(v.Eye, v.Center, v.Up).Perspective(v.Fovy, v.Width/v.Height, v.Near, v.Far)
	d(im, fauxgl.Matrix(m), fauxgl.Vector(v.Eye))
}

var mesh [3]*fauxgl.Mesh

func redraw(im *duitdraw.Image, matrix fauxgl.Matrix, eye fauxgl.Vector) {
	width, height := im.R.Dx(), im.R.Dy()

	// TODO: this allocates a new image each time.

	context := fauxgl.NewContext(width, height)
	context.ClearColor = fauxgl.Black
	context.ClearColorBuffer()

	light := fauxgl.V(-2, 0, 1).Normalize()
	colors := []fauxgl.Color{
		{0.5, 1, 0.65, 1},
		{1, 0.4, 0.5, 1},
		{0.3, 0.2, 1, 1},
	}

	for i, m := range mesh {
		m.MoveTo(fauxgl.V(float64(i), 0, 0), fauxgl.V(1, 1, 1))
		shader := fauxgl.NewPhongShader(matrix, light, eye)
		shader.ObjectColor = colors[i]
		context.Shader = shader
		context.DrawMesh(m)
	}

	im.DrawImage(im.R, context.Image(), image.ZP, imdraw.Src)
}

func position(t float64) (float64, float64) {
	// Position in the air.
	W, H := 1.0, 1.0
	x := W * t / 2
	y := 2*H*t - H*t*t
	// TODO...
	return x, y
}
