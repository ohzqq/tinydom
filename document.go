package tinydom

import (
	"syscall/js"
)

// Document wraps the JavaScript document element, which is usually fetched by js.Global().Get("document")
type Document struct {
	*BaseNode
}

var doc = js.Global().Get("document")

func GetDocument() *Document {
	return &Document{WrapNode(doc)}
}

func (e *Document) ActiveElement() *Element {
	return WrapElement(e.Get("activeElement"))
}

func (e *Document) Body() *Element {
	return WrapElement(e.Get("body"))
}

func (e *Document) Head() *Element {
	return WrapElement(e.Get("head"))
}

func (e *Document) DocumentElement() *Element {
	return WrapElement(e.Get("documentElement"))
}

func (d *Document) CreateElement(tag string) *Element {
	return WrapElement(d.Call("createElement", tag))
}

func (d *Document) CreateTextElement(textContent string) *Element {
	return WrapElement(d.Call("createTextNode", textContent))
}

func (d *Document) CreateDocumentFragment() *Element {
	return WrapElement(d.Call("createDocumentFragment"))
}

func (d *Document) DocumentURI() string {
	return d.Get("documentURI").String()
}

func (d *Document) FullscreenElement() *Element {
	return WrapElement(d.Get("fullscreenElement"))
}

func (d *Document) ExitFullscreen() {
	d.Call("exitFullscreen")
}

func (d *Document) GetElementById(id string) *Element {
	return WrapElement(d.Call("getElementById", id))
}

func (e *Document) QuerySelector(selectors string) *Element {
	return querySelector(e.Value, selectors)
}

func (e *Document) QuerySelectorAll(selectors string) []*Element {
	return querySelectorAll(e.Value, selectors)
}

func (e *Document) GetElementsByTagName(tagName string) []*Element {
	return getElementsByTagName(e.Value, tagName)
}

func (d *Document) Write(markup string) {
	d.Call("write", markup)
}

// DocumentFragment is the dom DocumentFragement interface.
type DocumentFragment struct {
	*BaseNode
}

// WrapDocumentFragment wraps a js.Value to *DocumentFragment.
func WrapDocumentFragment(val js.Value) *DocumentFragment {
	return &DocumentFragment{WrapNode(val)}
}

func (d *DocumentFragment) Append(nodes ...Node) {
	for _, node := range nodes {
		d.Call("append", node.Underlying())
	}
}

func (d *DocumentFragment) Prepend(nodes ...Node) {
	for _, node := range nodes {
		d.Call("prepend", node.Underlying())
	}
}

func (d *DocumentFragment) ChildElementCount() int {
	return d.Get("childElementCount").Int()
}

func (d *DocumentFragment) GetElementById(id string) *Element {
	return WrapElement(d.Call("getElementById", id))
}

func (d *DocumentFragment) FirstElementChild() *Element {
	return WrapElement(d.Call("firstElementChild"))
}

func (d *DocumentFragment) LastElementChild() *Element {
	return WrapElement(d.Call("lastElementChild"))
}

func (e *DocumentFragment) QuerySelector(selectors string) *Element {
	return querySelector(e.Value, selectors)
}

func (e *DocumentFragment) QuerySelectorAll(selectors string) []*Element {
	return querySelectorAll(e.Value, selectors)
}

func (e *DocumentFragment) Children() *HTMLCollection {
	return NewHTMLCollection(e)
}
