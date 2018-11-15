package main

import (
	"log"

	"github.com/fogleman/fauxgl"
	"github.com/ktye/duit"
	"github.com/ktye/duit/scene"
)

func main() {

	for i := 0; i < 3; i++ {
		mesh[i] = fauxgl.NewSphere(3)
		mesh[i].SmoothNormals()
		mesh[i].Transform(fauxgl.Scale(fauxgl.V(0.1, 0.1, 0.1)))
	}
	//mesh = fauxgl.NewCube()

	view := scene.NewView()

	dui, err := duit.NewDUI("", nil)
	if err != nil {
		panic(err)
	}

	dui.Top.UI = &scene.Scene{
		View:        view,
		SceneDrawer: draw(redraw),
	}
	dui.Render()

	for {
		select {
		case e := <-dui.Inputs:
			dui.Input(e)
		case err, ok := <-dui.Error:
			if !ok {
				return
			}
			log.Print(err)
		}
	}
}
