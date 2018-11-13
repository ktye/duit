package main

import (
	"log"

	"github.com/ktye/duit"
	"github.com/ktye/duit/scene"
)

func main() {
	view := scene.NewView()
	//view.Eye = scene.V(-1, -2, 2)
	view.Eye = scene.V(-10, -10, 2)
	view.Up = scene.V(0, 0, 1)
	view.Near = 0.1
	view.Far = 50
	view.Fovy = 20

	dui, err := duit.NewDUI("", nil)
	if err != nil {
		panic(err)
	}

	dui.Top.UI = &scene.Scene{
		View:        view,
		SceneDrawer: draw(render),
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
