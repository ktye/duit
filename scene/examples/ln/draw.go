package main

import (
	"image"
	imdraw "image/draw"

	"github.com/fogleman/gg"
	"github.com/fogleman/ln/ln"
	"github.com/ktye/duit/scene"
	"github.com/ktye/duitdraw"
)

type draw func(im *duitdraw.Image, eye, center, up ln.Vector, fovy, near, far, width, height float64)

func (d draw) DrawScene(im *duitdraw.Image, view scene.View) {
	eye := ln.Vector(view.Eye)
	center := ln.Vector(view.Center)
	up := ln.Vector(view.Up)
	fovy := view.Fovy
	near, far := view.Near, view.Far
	width, height := view.Width, view.Height

	d(im, eye, center, up, fovy, near, far, width, height)
}

func render(im *duitdraw.Image, eye, center, up ln.Vector, fovy, near, far, width, height float64) {
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{0, 0, 0}, ln.Vector{1, 1, 1}))
	scene.Add(ln.NewCube(ln.Vector{-0.5, -0.5, -0.5}, ln.Vector{0, 0, 0}))
	paths := scene.Render(eye, center, up, width, height, fovy, near, far, 0.01)
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
