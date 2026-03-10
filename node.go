//go:build js && wasm

package tinydom

import (
	"syscall/js"
)

type Node interface {
	Underlying() js.Value
}

type BaseNode struct {
	Node
	js.Value
}

func WrapNode(val js.Value) *BaseNode {
	return &BaseNode{Value: val}
}

func (b *BaseNode) Underlying() js.Value {
	return b.Value
}

func (e *BaseNode) AppendChild(child Node) {
	e.Call("appendChild", child.Underlying())
}

func (e *BaseNode) AppendChildren(children ...*BaseNode) {
	for _, child := range children {
		e.AppendChild(child)
	}
}

func (e *BaseNode) RemoveAllChildNodes() {
	for e.HasChildNodes() {
		e.RemoveChild(e.LastChild())
	}
}

func (e *BaseNode) ChildNodes() []*BaseNode {
	nodeList := e.Get("childNodes")
	length := nodeList.Get("length").Int()
	var nodes []*BaseNode
	for i := 0; i < length; i++ {
		nodes = append(nodes, WrapNode(nodeList.Call("item", i)))
	}
	return nodes
}

func (e *BaseNode) FirstChild() *BaseNode {
	return WrapNode(e.Get("firstChild"))
}

func (e *BaseNode) LastChild() *BaseNode {
	return WrapNode(e.Get("lastChild"))
}

func (e *BaseNode) NextSibling() *BaseNode {
	return WrapNode(e.Get("nextSibling"))
}

func (e *BaseNode) NodeType() int {
	return e.Get("nodeType").Int()
}

func (e *BaseNode) NodeValue() string {
	return e.Get("nodeValue").String()
}

func (e *BaseNode) SetNodeValue(s string) *BaseNode {
	e.Set("nodeValue", s)
	return e
}
func (e *BaseNode) ParentNode() *BaseNode {
	return WrapNode(e.Get("parentNode"))
}

func (e *BaseNode) TextContent() string {
	return e.Get("textContent").String()
}

func (e *BaseNode) SetTextContent(s string) *BaseNode {
	e.Set("textContent", s)
	return e
}

func (e *BaseNode) Contains(n Node) bool {
	return e.Call("contains", n.Underlying()).Bool()
}

func (e *BaseNode) HasChildNodes() bool {
	return e.Call("hasChildNodes").Bool()
}

func (e *BaseNode) InsertBefore(newNode, referenceNode Node) *BaseNode {
	return WrapNode(e.Call("insertBefore", newNode.Underlying(), referenceNode.Underlying()))
}

func (e *BaseNode) IsEqualNode(n Node) bool {
	return e.Call("isEqualNode", n.Underlying()).Bool()
}

func (e *BaseNode) IsSameNode(n Node) bool {
	return e.Call("isSameNode", n.Underlying()).Bool()
}

func (e *BaseNode) LookupPrefix() string {
	return e.Call("lookupPrefix").String()
}

func (e *BaseNode) Normalize() {
	e.Call("normalize")
}

func (e *BaseNode) RemoveChild(c Node) *BaseNode {
	return WrapNode(e.Call("removeChild", c.Underlying()))
}

func (e *BaseNode) ReplaceChild(newChild, oldChild Node) *BaseNode {
	return WrapNode(e.Call("replaceChild", newChild.Underlying(), oldChild.Underlying()))
}

func (e *BaseNode) AddEventListener(t string, listener js.Func) *BaseNode {
	e.Call("addEventListener", t, listener)
	return e
}

func (e *BaseNode) RemoveEventListener(t string, listener js.Func) *BaseNode {
	e.Call("removeEventListener", t, listener)
	return e
}

func (e *BaseNode) DispatchEvent(event *Event) *BaseNode {
	e.Call("dispatchEvent", event.Value)
	return e
}

func querySelector(e js.Value, selectors string) *Element {
	return WrapElement(e.Call("querySelector", selectors))
}

func querySelectorAll(e js.Value, selectors string) []*Element {
	nodeList := e.Call("querySelectorAll", selectors)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = WrapElement(nodeList.Call("item", i))
	}

	return nodes
}

func getElementsByTagName(e js.Value, tagName string) []*Element {
	nodeList := e.Call("getElementsByTagName", tagName)
	length := nodeList.Get("length").Int()

	nodes := make([]*Element, length)

	for i := 0; i < length; i++ {
		nodes[i] = WrapElement(nodeList.Call("item", i))
	}

	return nodes
}
