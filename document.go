//go:build js && wasm

package tinydom

import "syscall/js"

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
