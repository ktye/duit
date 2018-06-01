package main

import (
	"bytes"
	"image"
	"log"
	"os/exec"
	"strings"

	"github.com/ktye/duit"
)

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s\n", msg, err)
	}
}

func main() {
	dui, err := duit.NewDUI("ex/fonts", nil)
	check(err, "new dui")

	buf, err := exec.Command("fontsrv", "-p", ".").Output()
	check(err, "listing fonts using fontsrv")
	fonts := strings.Split(string(buf), "\n")

	fontValues := make([]*duit.ListValue, len(fonts))
	for i, s := range fonts {
		fontValues[i] = &duit.ListValue{Text: s}
	}

	src := bytes.NewReader([]byte(`0 1 2 3 4 5 6 7 8 9
a b c d e f g h i j k l m n o p q r s t u v w x y z`))
	edit, err := duit.NewEdit(src)
	check(err, "new edit")

	var fontList *duit.List
	fontList = &duit.List{
		Values: fontValues,
		Changed: func(index int) (e duit.Event) {
			lv := fontList.Values[index]
			// xxx todo should free font, but that seems to hang draw
			if lv.Selected {
				go func() {
					font, err := dui.Display.OpenFont("/mnt/font/" + lv.Text + "15a/font")
					check(err, "open font")
					dui.Call <- func() {
						edit.Font = font
						dui.MarkDraw(edit)
						dui.Draw()
					}
				}()
			} else {
				edit.Font = nil
			}
			e.NeedDraw = true
			return
		},
	}

	search := &duit.Field{
		Placeholder: "search...",
		Changed: func(s string) (e duit.Event) {
			s = strings.ToLower(s)
			nl := []*duit.ListValue{}
			for _, lv := range fontValues {
				if lv.Selected || strings.Contains(strings.ToLower(lv.Text), s) {
					nl = append(nl, lv)
				}
			}
			fontList.Values = nl
			e.NeedDraw = true
			return
		},
	}

	dui.Top.UI = &duit.Split{
		Gutter: 1,
		Split: func(width int) []int {
			first := dui.Scale(250)
			return []int{first, width - first}
		},
		Kids: duit.NewKids(
			&duit.Box{
				Padding: duit.SpaceXY(6, 4),
				Margin:  image.Pt(0, 4),
				Kids: duit.NewKids(
					search,
					duit.NewScroll(fontList),
				),
			},
			edit,
		),
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
