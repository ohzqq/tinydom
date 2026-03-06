package tinydom

import "syscall/js"

// Document wraps the JavaScript document element, which is usually fetched by js.Global().Get("document")
type Document struct {
	js.Value
}

var doc = js.Global().Get("document")

func GetDocument() *Document {
	return &Document{doc}
}

func (e *Document) ActiveElement() *Element {
	return &Element{e.Get("activeElement")}
}

func (e *Document) DocumentElement() *Element {
	return &Element{e.Get("documentElement")}
}

func (d *Document) CreateElement(tag string) *Element {
	return &Element{d.Call("createElement", tag)}
}

func (d *Document) CreateTextNode(textContent string) *Element {
	return &Element{d.Call("createTextNode", textContent)}
}

func (d *Document) CreateDocumentFragment() *Element {
	return &Element{d.Call("createDocumentFragment")}
}

func (d *Document) GetElementById(id string) *Element {
	return &Element{d.Call("getElementById", id)}
}

func (e *Document) GetElementsByTagName(tagName string) []*Element {
	nodeList := e.Call("getElementsByTagName", tagName)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = &Element{nodeList.Call("item", i)}
	}

	return nodes
}

func (e *Document) QuerySelector(selectors string) *Element {
	return &Element{e.Call("querySelector", selectors)}
}

func (e *Document) QuerySelectorAll(selectors string) []*Element {
	nodeList := e.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = &Element{nodeList.Call("item", i)}
	}

	return nodes
}

func (d *Document) Write(markup string) {
	d.Call("write", markup)
}
