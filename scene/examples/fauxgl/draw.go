package main

import (
	"image"
	imdraw "image/draw"

	"github.com/fogleman/fauxgl"
	"github.com/ktye/duit/scene"
	"github.com/ktye/duitdraw"
)

// Draw stores the redraw function.
// It is typed to draw, which implements a scene.SceneDrawer.
//var Draw draw = redraw

// Draw implements a scene.SceneDrawer.
// It has the same signature as the redraw function, which is converted to a SceneDrawer by type casting.
type draw func(im *duitdraw.Image, mat fauxgl.Matrix, eye fauxgl.Vector)

func (d draw) DrawScene(im *duitdraw.Image, view scene.View) {
	d(im, fauxgl.Matrix(view.Matrix()), fauxgl.Vector(view.Eye))
}

var mesh *fauxgl.Mesh

func redraw(im *duitdraw.Image, matrix fauxgl.Matrix, eye fauxgl.Vector) {
	if mesh == nil {
		if m, err := fauxgl.LoadSTL("hello.stl"); err != nil {
			panic(err)
		} else {
			mesh = m
		}
		mesh.BiUnitCube()
		mesh.SmoothNormalsThreshold(fauxgl.Radians(30))
	}
	width, height := im.R.Dx(), im.R.Dy()

	// TODO: this allocates a new image each time.

	context := fauxgl.NewContext(width, height)
	context.ClearColor = fauxgl.Black
	context.ClearColorBuffer()

	light := fauxgl.V(-2, 0, 1).Normalize()
	color := fauxgl.Color{0.5, 1, 0.65, 1}

	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader
	context.DrawMesh(mesh)

	im.DrawImage(im.R, context.Image(), image.ZP, imdraw.Src)
}
