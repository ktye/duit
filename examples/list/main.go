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
	dui, err := duit.NewDUI("ex/list", nil)
	check(err, "new dui")

	dui.Top.UI = &duit.List{
		Values: []*duit.ListValue{
			{Text: "item 1"},
			{Text: "item 2"},
			{Text: "item 3"},
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
