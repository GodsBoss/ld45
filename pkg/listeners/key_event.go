package listeners

import (
	"github.com/gopherjs/gopherjs/js"
)

func WithKeyEvent(f func(event KeyEvent)) func(*js.Object, []*js.Object) interface{} {
	return func(_ *js.Object, arguments []*js.Object) interface{} {
		f(
			KeyEvent{
				obj: arguments[0],
			},
		)
		return nil
	}
}

type KeyEvent struct {
	obj *js.Object
}

func (event KeyEvent) CharCode() int {
	return event.obj.Get("charCode").Int()
}

func (event KeyEvent) Code() string {
	return event.obj.Get("code").String()
}

func (event KeyEvent) Key() string {
	return event.obj.Get("key").String()
}

func (event KeyEvent) KeyCode() int {
	return event.obj.Get("keyCode").Int()
}

func (event KeyEvent) AltKey() bool {
	return event.obj.Get("altKey").Bool()
}

func (event KeyEvent) CtrlKey() bool {
	return event.obj.Get("ctrlKey").Bool()
}

func (event KeyEvent) ShiftKey() bool {
	return event.obj.Get("shiftkey").Bool()
}

func (event KeyEvent) Type() string {
	return event.obj.Get("type").String()
}

func (event KeyEvent) Repeat() bool {
	return event.obj.Get("repeat").Bool()
}
