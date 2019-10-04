package dom

import (
	"github.com/gopherjs/gopherjs/js"
)

// Node is an HTML node.
type Node interface {
	exposedObject
	isNode() bool
}

type exposedObject interface {
	exposeObject() *js.Object
}

// Element represents a generic HTML element. It implements Node.
type Element struct {
	obj *js.Object
}

func (element *Element) isNode() bool {
	return true
}

func (element *Element) exposeObject() *js.Object {
	return element.obj
}

// AppendChild appends a node to the element.
func (element *Element) AppendChild(node Node) {
	element.obj.Call("appendChild", node.exposeObject())
}
