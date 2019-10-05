package date

import (
	"github.com/gopherjs/gopherjs/js"
)

type Date struct {
	obj *js.Object
}

func Now() Date {
	return Date{
		obj: js.Global.Get("Date").New(),
	}
}

func FromUnixTimestamp(timestamp int) Date {
	return Date{
		obj: js.Global.Get("Date").New(timestamp),
	}
}

func FromString(str string) Date {
	return Date{
		obj: js.Global.Get("Date").New(str),
	}
}

func (date Date) Unix() int {
	return date.obj.Call("valueOf").Int()
}

func (date Date) Add(ms int) Date {
	return FromUnixTimestamp(date.Unix() + ms)
}

func (date Date) Sub(otherDate Date) int {
	return date.Unix() - otherDate.Unix()
}

func (date Date) Milliseconds() int {
	return date.obj.Call("getMilliseconds").Int()
}

func (date Date) Seconds() int {
	return date.obj.Call("getSeconds").Int()
}

func (date Date) Minutes() int {
	return date.obj.Call("getMinutes").Int()
}

func (date Date) Hours() int {
	return date.obj.Call("getHours").Int()
}
