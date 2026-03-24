//go:build js && wasm

package tinydom

import "syscall/js"

type Event struct {
	js.Value
}

func NewEvent(e string, opts ...*EventOptions) *Event {
	return newEvent(e, false, opts...)
}

func NewCustomEvent(e string, opts ...*EventOptions) *Event {
	return newEvent(e, true, opts...)
}

func WrapEvent(e js.Value) *Event {
	return &Event{Value: e}
}

func newEvent(e string, custom bool, opts ...*EventOptions) *Event {
	event := "Event"
	if custom {
		event = "CustomEvent"
	}
	args := []any{e}
	if len(opts) > 0 {
		args = append(args, opts[0].Value)
	}
	return &Event{Value: js.Global().Get(event).New(args...)}
}

func (e *Event) Target() *Element {
	return WrapElement(e.Get("target"))
}

func (e *Event) PreventDefault() {
	e.Call("preventDefault")
}

func (e *Event) StopImmediatePropagation() {
	e.Call("stopImmediatePropagation")
}

func (e *Event) StopPropagation() {
	e.Call("stopPropagation")
}

func (e *Event) Code() string {
	return e.Get("code").String()
}

func (e *Event) Key() string {
	return e.Get("key").String()
}

func (e *Event) Type() string {
	return e.Get("type").String()
}

func (e *Event) KeyCode() int {
	return e.Get("keyCode").Int()
}

func (e *Event) Detail() js.Value {
	return e.Get("detail")
}

func (e *Event) Bubbles() bool {
	return e.Get("bubbles").Bool()
}

func (e *Event) Composed() bool {
	return e.Get("composed").Bool()
}

type EventOptions struct {
	js.Value
}

func NewEventOptions() *EventOptions {
	return &EventOptions{Value: js.ValueOf(map[string]any{})}
}

func (e *EventOptions) SetBubbles() *EventOptions {
	e.Set("bubbles", true)
	return e
}

func (e *EventOptions) SetCancelable() *EventOptions {
	e.Set("cancelable", true)
	return e
}

func (e *EventOptions) SetComposed() *EventOptions {
	e.Set("composed", true)
	return e
}

func (e *EventOptions) SetDetails(details js.Value) *EventOptions {
	e.Set("detail", details)
	return e
}
