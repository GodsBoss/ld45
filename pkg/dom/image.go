package dom

import (
	"github.com/GodsBoss/ld45/pkg/listeners"

	"github.com/gopherjs/gopherjs/js"
)

type Image struct {
	obj *js.Object
}

func (img *Image) Src(url string) *Image {
	img.obj.Set("src", url)
	return img
}

func (img *Image) OnLoad(onSuccess func()) *Image {
	listeners.Add(img.obj, "load", listeners.Simple(onSuccess))
	return img
}

func (img *Image) OnError(onError func(listeners.Event)) *Image {
	listeners.Add(img.obj, "error", listeners.WithEvent(onError))
	return img
}
