package tinydom

import (
	"syscall/js"
)

type Node struct {
	js.Value
}

func WrapNode(val js.Value) *Node {
	return &Node{Value: val}
}

func (e *Node) AppendChild(child *Node) {
	e.Call("appendChild", child.Value)
}

func (e *Node) AppendChildren(children ...*Node) {
	for _, child := range children {
		e.AppendChild(child)
	}
}

func (e *Node) RemoveAllChildNodes() {
	for e.HasChildNodes() {
		e.RemoveChild(e.LastChild())
	}
}

func (e *Node) ChildNodes() []*Node {
	nodeList := e.Get("childNodes")
	length := nodeList.Get("length").Int()
	var nodes []*Node
	for i := 0; i < length; i++ {
		nodes = append(nodes, &Node{nodeList.Call("item", i)})
	}
	return nodes
}

func (e *Node) FindChildNode(tag string) *Node {
	children := e.ChildNodes()
	for _, child := range children {
		if child.TagName() == tag {
			return child
		}
	}

	return nil
}

func (e *Node) FirstChild() *Node {
	return &Node{e.Get("firstChild")}
}

func (e *Node) LastChild() *Node {
	return &Node{e.Get("lastChild")}
}

func (e *Node) NextSibling() *Node {
	return &Node{e.Get("nextSibling")}
}

func (e *Node) NodeType() int {
	return e.Get("nodeType").Int()
}

func (e *Node) NodeValue() string {
	return e.Get("nodeValue").String()
}

func (e *Node) SetNodeValue(s string) *Node {
	e.Set("nodeValue", s)
	return e
}
func (e *Node) ParentNode() *Node {
	return &Node{e.Get("parentNode")}
}

func (e *Node) TextContent() string {
	return e.Get("textContent").String()
}

func (e *Node) SetTextContent(s string) *Node {
	e.Set("textContent", s)
	return e
}

func (e *Node) Contains(n *Node) bool {
	return e.Call("contains", n).Bool()
}

func (e *Node) HasChildNodes() bool {
	return e.Call("hasChildNodes").Bool()
}

func (e *Node) InsertBefore(newNode, referenceNode *Node) *Node {
	return &Node{e.Call("insertBefore", newNode, referenceNode)}
}

func (e *Node) IsEqualNode(n *Node) bool {
	return e.Call("isEqualNode", n).Bool()
}

func (e *Node) IsSameNode(n *Node) bool {
	return e.Call("isSameNode", n).Bool()
}

func (e *Node) LookupPrefix() string {
	return e.Call("lookupPrefix").String()
}

func (e *Node) Normalize() {
	e.Call("normalize")
}

func (e *Node) RemoveChild(c *Node) *Node {
	return &Node{e.Call("removeChild", c)}
}

func (e *Node) ReplaceChild(newChild, oldChild *Node) *Node {
	return &Node{e.Call("replaceChild", newChild, oldChild)}
}

func (e *Node) AddEventListener(t string, listener js.Func) *Element {
	e.Call("addEventListener", t, listener)
	return e
}

func (e *Node) RemoveEventListener(t string, listener js.Func) *Element {
	e.Call("removeEventListener", t, listener)
	return e
}
