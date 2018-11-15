package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fogleman/fauxgl"
	"github.com/ktye/duit"
	"github.com/ktye/duit/scene"
)

func main() {

	if len(os.Args) == 2 {
		loadmesh(os.Args[1])
	} else {
		loadmesh("fauxgl/examples/hello.stl")
	}

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

// Load the file from $GOPATH/src/github.com/fogleman/$file.
func loadmesh(file string) {
	p := filepath.Join(os.Getenv("GOPATH"), "src/github.com/fogleman", file)
	if m, err := fauxgl.LoadSTL(p); err != nil {
		fmt.Println("could not load file from: ", p)
		panic(err)
	} else {
		mesh = m
	}
	mesh.BiUnitCube()
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))
}
