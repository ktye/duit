package main

import (
	"log"

	"github.com/ktye/duit"
	"github.com/ktye/duit/scene"
)

func main() {

	view := scene.View{
		Eye:    scene.V(-1, -2, 2),
		Center: scene.V(0, 0, 0),
		Up:     scene.V(0, 0, 1),
		Near:   1,
		Far:    50,
		Fovy:   20,
	}

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
