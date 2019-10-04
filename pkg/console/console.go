package console

import (
	"github.com/gopherjs/gopherjs/js"
)

var c = js.Global.Get("console")

func Log(args ...interface{}) {
	c.Call("log", args...)
}
