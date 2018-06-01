package main

import (
	"fmt"
	"log"

	"github.com/ktye/duit"

	draw "github.com/ktye/duitdraw"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func main() {
	dui, err := duit.NewDUI("ex/gridlist", nil)
	check(err, "new dui")

	var rows []*duit.Gridrow
	for i := 0; i < 30; i++ {
		values := []string{
			fmt.Sprintf("cell 0,%d", i),
			fmt.Sprintf("cell 1,%d", i),
			fmt.Sprintf("cell 2,%d", i),
		}
		row := &duit.Gridrow{
			Selected: i%10 == 0,
			Values:   values,
		}
		rows = append(rows, row)
	}
	rows = append([]*duit.Gridrow{&duit.Gridrow{
		Values: []string{
			"and this is the longest of them all! and this is the longest of them all! and this is the longest of them all!",
			"this is quite a long line",
			"but this is is even longer",
		},
	}}, rows...)

	dui.Top.UI = duit.NewScroll(
		&duit.Gridlist{
			Header:   &duit.Gridrow{Values: []string{"col1", "col2", "col3"}},
			Rows:     rows,
			Multiple: true,
			Striped:  true,
			Padding:  duit.SpaceXY(10, 2),
			Halign: []duit.Halign{
				duit.HalignMiddle,
				duit.HalignRight,
				duit.HalignLeft,
			},
			Changed: func(index int) (e duit.Event) {
				log.Printf("gridlist, index %d changed\n", index)
				return
			},
			Click: func(index int, m draw.Mouse) (e duit.Event) {
				log.Printf("gridlist, click, index %d, m %d\n", index, m)
				return
			},
			Keys: func(k rune, m draw.Mouse) (e duit.Event) {
				log.Printf("gridlist, key %c, mouse %v\n", k, m)
				return
			},
		},
	)
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
