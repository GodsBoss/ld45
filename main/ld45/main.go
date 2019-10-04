package main

import (
	"github.com/GodsBoss/ld45/pkg/listeners"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	listeners.Add(js.Global.Get("window"), "load", initialize)
}

func initialize(this *js.Object, arguments []*js.Object) interface{} {
	return nil
}
