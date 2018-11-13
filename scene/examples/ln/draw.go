package main

import (
	"image"
	imdraw "image/draw"

	"unsafe"

	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"
	"github.com/ktye/duit/scene"
	"github.com/ktye/duitdraw"
)

type draw func(im *duitdraw.Image, matrix ln.Matrix, eye ln.Vector, width, height float64)

func (d draw) DrawScene(im *duitdraw.Image, view scene.View) {
	eye := ln.Vector(view.Eye)

	// ln.Matrix elements are not exported.
	// We need to set it from a scene.Matrix.
	mat := view.Matrix()
	var m ln.Matrix = *(*ln.Matrix)(unsafe.Pointer(&mat))

	w, h := view.Width, view.Height
	d(im, m, eye, w, h)
}

func render(im *duitdraw.Image, matrix ln.Matrix, eye ln.Vector, width, height float64) {
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{0, 0, 0}, ln.Vector{1, 1, 1}))
	scene.Add(ln.NewCube(ln.Vector{-0.5, -0.5, -0.5}, ln.Vector{0, 0, 0}))
	paths := scene.RenderWithMatrix(matrix, eye, width, height, 0.01)
	img := pathsToImage(paths, width, height)
	im.DrawImage(im.R, img, image.ZP, imdraw.Src)
}

func pathsToImage(paths ln.Paths, width, height float64) image.Image {
	scale := 1.0
	w, h := int(width*scale), int(height*scale)
	dc := gg.NewContext(w, h)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	for _, path := range paths {
		for _, v := range path {
			dc.LineTo(v.X*scale, v.Y*scale)
		}
		dc.NewSubPath()
	}
	dc.Stroke()
	return dc.Image()
}
