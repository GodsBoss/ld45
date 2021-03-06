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

func (canvas *Canvas) Width() int {
	return canvas.obj.Get("width").Int()
}

func (canvas *Canvas) Height() int {
	return canvas.obj.Get("height").Int()
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

func (ctx *Context2D) Width() int {
	return ctx.canvas.Width()
}

func (ctx *Context2D) Height() int {
	return ctx.canvas.Height()
}

// EnableImageSmoothing enables image smoothing. This is also the default.
func (ctx *Context2D) EnableImageSmoothing() *Context2D {
	return ctx.xableImageSmoothing(true)
}

// DisableImageSmoothing disbles image smoothing, good for pixel art.
func (ctx *Context2D) DisableImageSmoothing() *Context2D {
	return ctx.xableImageSmoothing(false)
}

func (ctx *Context2D) xableImageSmoothing(value bool) *Context2D {
	ctx.obj.Set("imageSmoothingEnabled", value)
	return ctx
}

type FillStyle interface {
	applyTo(ctx *Context2D)
}

func Color(colorString string) FillStyle {
	return colorFillStyle(colorString)
}

type colorFillStyle string

func (style colorFillStyle) applyTo(ctx *Context2D) {
	ctx.obj.Set("fillStyle", string(style))
}

func (ctx *Context2D) FillStyle(fillStyle FillStyle) *Context2D {
	fillStyle.applyTo(ctx)
	return ctx
}

func (ctx *Context2D) ClearRect(x, y, w, h float64) *Context2D {
	ctx.obj.Call("clearRect", x, y, w, h)
	return ctx
}

func (ctx *Context2D) FillRect(x, y, w, h float64) *Context2D {
	ctx.obj.Call("fillRect", x, y, w, h)
	return ctx
}

type ImageSource interface {
	source() *js.Object
}

type imageSource struct {
	src *js.Object
}

func (src imageSource) source() *js.Object {
	return src.src
}

func ImageElementSource(img *Image) ImageSource {
	return imageSource{
		src: img.obj,
	}
}

func (ctx *Context2D) DrawImage(source ImageSource, sourceX, sourceY, sourceWidth, sourceHeight, destinationX, destinationY, destinationWidth, destinationHeight int) *Context2D {
	ctx.obj.Call(
		"drawImage",
		source.source(),
		sourceX, sourceY,
		sourceWidth, sourceHeight,
		destinationX, destinationY,
		destinationWidth, destinationHeight,
	)
	return ctx
}
