package main

import (
	"log"

	"github.com/ktye/duit"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func main() {
	dui, err := duit.NewDUI("ex/button", nil)
	check(err, "new dui")

	dui.Top.UI = &duit.Button{
		Text: "click me",
		Click: func() (e duit.Event) {
			log.Printf("clicked\n")
			return
		},
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
			log.Printf("duit: %s\n", err)
		}
	}
}
