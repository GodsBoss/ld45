package dom

import (
	"github.com/GodsBoss/ld45/pkg/listeners"

	"github.com/gopherjs/gopherjs/js"
)

type Window struct {
	obj *js.Object
}

func GlobalWindow() *Window {
	return &Window{
		obj: js.Global.Get("window"),
	}
}

func (window *Window) SetTimeout(f func(), delayInMS int) *Window {
	g := func() {
		go f()
	}
	window.obj.Call("setTimeout", g, delayInMS)
	return window
}

func (window *Window) OnKeyDown(f func(listeners.KeyEvent)) *Window {
	listeners.Add(window.obj, "keydown", listeners.WithKeyEvent(f))
	return window
}

func (window *Window) OnKeyUp(f func(listeners.KeyEvent)) *Window {
	listeners.Add(window.obj, "keyup", listeners.WithKeyEvent(f))
	return window
}

func (window *Window) OnKeyPress(f func(listeners.KeyEvent)) *Window {
	listeners.Add(window.obj, "keypress", listeners.WithKeyEvent(f))
	return window
}
