//go:build js && wasm

package tinydom

import (
	"errors"
	"strings"
	"syscall/js"
)

type Element struct {
	*BaseNode
}

func WrapElement(val js.Value) *Element {
	return &Element{BaseNode: WrapNode(val)}
}

func (e *Element) Underlying() js.Value {
	return e.BaseNode.Value
}

func (e *Element) HasFocus() bool {
	return e.IsEqualNode(WrapNode(GetDocument().ActiveElement().Value))
}

func (e *Element) AppendBefore(n *Element) {
	e.ParentNode().InsertBefore(WrapNode(n.Value), WrapNode(e.Value))
}

func (e *Element) AppendAfter(n *Element) {
	e.ParentNode().InsertBefore(WrapNode(n.Value), e.NextSibling())
}

func (e *Element) SetId(id string) *Element {
	e.Set("id", id)
	return e
}

func (e *Element) SetAttribute(key, value interface{}) *Element {
	e.Call("setAttribute", key, value)
	return e
}

func (e *Element) SetClass(values ...string) *Element {
	return e.SetMultiValueAttribute("class", values...)
}

// ErrClassAlreadyExisting is being thrown when trying to append the same class multiple times
var ErrClassAlreadyExisting = errors.New("tried to append class multiple times")

func (e *Element) AppendClass(values ...string) error {
	existing, currentClasses := e.Class()

	if !existing {
		e.SetClass(values...)
		return nil
	}

	for _, newClass := range values {
		for _, existingclass := range currentClasses {
			if newClass == existingclass {
				return ErrClassAlreadyExisting
			}
		}
	}

	newClass := append(currentClasses, values...)
	e.SetClass(newClass...)

	return nil
}

func (e *Element) Class() (bool, []string) {
	exists, attributeValues := e.GetAttribute("class")
	if !exists {
		return false, nil
	}

	splittedValues := strings.Split(attributeValues, " ")

	result := make([]string, len(attributeValues))
	for i, value := range splittedValues {
		result[i] = value
	}

	return true, result
}

func (e *Element) ClassList() *DOMTokenList {
	return &DOMTokenList{Value: e.Get("classList")}
}

func (e *Element) SetMultiValueAttribute(attributeName string, values ...string) *Element {
	var value string

	valueCount := len(values)

	for i, rel := range values {
		value += rel

		if i < valueCount {
			value += " "
		}
	}

	e.SetAttribute(attributeName, value)
	return e
}

// AppendChildBr appends the child and adds an additional br
func (e *Element) AppendChildBr(child *Element) {
	e.Call("appendChild", child)
	e.Call("appendChild", GetDocument().CreateElement("br"))
}

func (e *Element) AppendChildrenBr(children ...*Element) {
	for _, child := range children {
		e.AppendChildBr(child)
	}
}

func (e *Element) Br() {
	br := GetDocument().CreateElement("br")
	e.AppendChild(br)
}

func (e *Element) QuerySelector(selectors string) *Element {
	return querySelector(e.Value, selectors)
}

func (e *Element) QuerySelectorAll(selectors string) []*Element {
	return querySelectorAll(e.Value, selectors)
}

func (e *Element) GetElementsByTagName(tagName string) []*Element {
	return querySelectorAll(e.Value, tagName)
}

func (e *Element) SetInnerHTML(value string) *Element {
	e.Set("innerHTML", value)
	return e
}

func (e *Element) InnerHTML() string {
	return e.Get("innerHTML").String()
}

func (e *Element) OuterHTML() string {
	return e.Get("outerHTML").String()
}

func (e *Element) SetOuterHTML(html string) *Element {
	e.Set("outerHTML", html)
	return e
}

// GetAttribute returns the searched attribute, returns false if the attribute wasn't found.
func (e *Element) GetAttribute(name string) (bool, string) {
	if !e.HasAttribute(name) {
		return false, ""
	}

	return true, e.Call("getAttribute", name).String()
}

func (e *Element) HasAttribute(name string) bool {
	return e.Call("hasAttribute", name).Bool()
}

func (e *Element) FindChildNode(tag string) *Element {
	children := e.ChildNodes()
	for _, child := range children {
		c := WrapElement(child.Value)
		if c.TagName() == tag {
			return c
		}
	}

	return nil
}

func (e *Element) TagName() string {
	return e.Get("tagName").String()
}

func (e *Element) Name() string {
	return e.Get("name").String()
}

func (e *Element) SetName(n string) *Element {
	e.Set("name", n)
	return e
}

func (e *Element) Style() *CSS {
	return &CSS{e.Get("style")}
}

func (e *Element) Blur() *Element {
	e.Call("blur")
	return e
}

func (e *Element) Focus() *Element {
	e.Call("focus")
	return e
}

// Copyright (c) 2014 Dominik Honnef
// MIT License
// from  github.com/dominikh/go-js-dom

func (e *Element) Dataset() map[string]string {
	o := e.Get("dataset")
	data := map[string]string{}
	keys := jsKeys(o)
	for _, key := range keys {
		data[key] = o.Get(key).String()
	}
	return data
}

// jsKeys returns the keys of the given JavaScript object.
func jsKeys(o js.Value) []string {
	if o.IsNull() || o.IsUndefined() {
		return nil
	}
	a := js.Global().Get("Object").Call("keys", o)
	s := make([]string, a.Length())
	for i := 0; i < a.Length(); i++ {
		s[i] = a.Index(i).String()
	}
	return s
}
