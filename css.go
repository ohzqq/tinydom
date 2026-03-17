//go:build js && wasm

package tinydom

import "syscall/js"

type CSS struct {
	js.Value
}

func (s *CSS) SetProperty(name string, val string) {
	s.Call("setProperty", name, val)
}

func (s *CSS) GetProperty(name string) string {
	return s.Get(name).String()
}

func (s *CSS) GetPropertyPriority(name string) string {
	return s.Call("getPropertyPriority", name).String()
}

func (s *CSS) CssText() string {
	return s.Get("cssText").String()
}

func (s *CSS) Item(idx int) string {
	return s.Call("item", idx).String()
}

func (s *CSS) Length() int {
	return s.Get("length").Int()
}

func (s *CSS) RemoveProperty(name string) {
	s.Call("removeProperty", name)
}
