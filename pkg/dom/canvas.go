package dom

import (
	"github.com/gopherjs/gopherjs/js"
)

// Canvas is an HTML canvas element. It implements Node.
type Canvas struct {
	obj *js.Object
}

func (canvas *Canvas) isNode() bool {
	return true
}

func (canvas *Canvas) exposeObject() *js.Object {
	return canvas.obj
}

// Resize resizes the canvas.
func (canvas *Canvas) Resize(width, height int) *Canvas {
	canvas.obj.Set("width", width)
	canvas.obj.Set("height", height)
	return canvas
}

// GetContext2D creates a 2D context for a canvas.
func (canvas *Canvas) GetContext2D() *Context2D {
	obj := canvas.obj.Call("getContext", "2d")
	return &Context2D{
		obj:    obj,
		canvas: canvas,
	}
}

type Context2D struct {
	obj    *js.Object
	canvas *Canvas
}
