package listeners

import (
	"github.com/gopherjs/gopherjs/js"
)

// Add adds an event listener to the given object.
func Add(
	obj *js.Object,
	eventType string,
	f func(
		this *js.Object,
		arguments []*js.Object,
	) interface{},
) *js.Object {
	g := func(this *js.Object, arguments []*js.Object) interface{} {
		// Non-blocking, as recommended by https://github.com/gopherjs/gopherjs#goroutines
		go f(this, arguments)
		return nil
	}
	return obj.Call("addEventListener", eventType, js.MakeFunc(g), false)
}

// Simple wraps a function which does not need this or arguments when used as an event handler. It returns a function compatible with Add().
func Simple(f func()) func(*js.Object, []*js.Object) interface{} {
	return func(*js.Object, []*js.Object) interface{} {
		f()
		return nil
	}
}

func WithEvent(f func(event Event)) func(*js.Object, []*js.Object) interface{} {
	return func(_ *js.Object, arguments []*js.Object) interface{} {
		f(
			Event{
				obj: arguments[0],
			},
		)
		return nil
	}
}

type Event struct {
	obj *js.Object
}

func (event Event) Type() string {
	return event.obj.Get("type").String()
}
