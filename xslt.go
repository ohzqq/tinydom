//go:build js && wasm

package tinydom

import "syscall/js"

type XSLTProcessor struct {
	js.Value
}

func NewXSLTProcessor() *XSLTProcessor {
	return &XSLTProcessor{GetWindow().Get("XSLTProcessor").New()}
}

func (x *XSLTProcessor) ImportStylesheet(node *Node) {
	x.Call("importStylesheet", node)
}

func (x *XSLTProcessor) TransformToFragment(node, document *Node) *Node {
	return &Node{x.Call("transformToFragment", node, document)}
}

func (x *XSLTProcessor) TransformToDocument(node *Node) *Node {
	return &Node{x.Call("transformToDocument", node)}
}
