//go:build js && wasm

package tinydom

import "syscall/js"

type XSLTProcessor struct {
	js.Value
}

func NewXSLTProcessor() *XSLTProcessor {
	return &XSLTProcessor{GetWindow().Get("XSLTProcessor").New()}
}

func (x *XSLTProcessor) ImportStylesheet(node Node) {
	x.Call("importStylesheet", node.Underlying())
}

func (x *XSLTProcessor) TransformToFragment(n, document Node) *BaseNode {
	return WrapNode(x.Call("transformToFragment", n.Underlying(), document.Underlying()))
}

func (x *XSLTProcessor) TransformToDocument(n Node) *BaseNode {
	return WrapNode(x.Call("transformToDocument", n.Underlying()))
}
