package main

import (
	"github.com/GodsBoss/ld45/pkg/console"
	"github.com/GodsBoss/ld45/pkg/dom"
	"github.com/GodsBoss/ld45/pkg/listeners"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	listeners.Add(js.Global.Get("window"), "load", initialize)
}

func initialize(this *js.Object, arguments []*js.Object) interface{} {
	doc, err := dom.GlobalDocument()
	if err != nil {
		console.Log("Initialization failed", err.Error())
		return nil
	}
	canvas := doc.CreateCanvas(800, 600)
	doc.Body().AppendChild(canvas)
	dom.SetStyles(
		canvas,
		map[string]string{
			"border":       "1px solid #000",
			"display":      "block",
			"margin-left":  "auto",
			"margin-right": "auto",
			"width":        "800px",
		},
	)
	return nil
}
