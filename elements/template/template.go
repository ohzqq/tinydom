package template

import (
	"syscall/js"

	"github.com/ohzqq/tinydom"
)

type Template struct {
	*tinydom.Element
}

func New() *Template {
	tmpl := tinydom.GetDocument().CreateElement("template")
	return &Template{
		Element: tmpl,
	}
}

func GetById(id string) *Template {
	return WrapTemplate(tinydom.GetDocument().GetElementById(id).Underlying())
}

func QuerySelector(sel string) *Template {
	return WrapTemplate{tinydom.GetDocument().QuerySelector(sel).Underlying()}
}

func WrapTemplate(val js.Value) *Template {
	return &Template{Element: tinydom.WrapElement(val)}
}

func (t *Template) ImportNode() *tinydom.DocumentFragment {
	content := tinydom.GetDocument().Call("importNode", t.Content().Underlying(), true)
	return tinydom.WrapDocumentFragment(content)
}

func (t *Template) Content() *tinydom.DocumentFragment {
	return tinydom.WrapDocumentFragment(t.Underlying().Get("content"))
}
