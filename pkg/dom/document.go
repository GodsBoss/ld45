package dom

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type Document struct {
	obj *js.Object
}

func (document *Document) Body() *Element {
	return &Element{
		obj: document.obj.Get("body"),
	}
}

func (document *Document) GetElementByID(id string) *Element {
	return &Element{
		obj: document.obj.Call("getElementById", id),
	}
}

func (document *Document) CreateCanvas(width, height int) *Canvas {
	canvas := &Canvas{
		obj: document.createElement("canvas"),
	}
	return canvas.Resize(width, height)
}

func (document *Document) CreateImage() *Image {
	return &Image{
		obj: document.createElement("img"),
	}
}

func (document *Document) createElement(tagName string) *js.Object {
	return document.obj.Call("createElement", tagName)
}

func GlobalDocument() (*Document, error) {
	obj := js.Global.Get("document")
	if !obj.Bool() {
		return nil, errors.New("document not found")
	}
	return &Document{
		obj: obj,
	}, nil
}
