package dom

import (
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
