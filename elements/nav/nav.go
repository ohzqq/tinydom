package nav

import (
	"github.com/ohzqq/tinydom"
	"github.com/ohzqq/tinydom/elements/li"
)

// Nav is a nav element
// See https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav for reference
type Nav struct {
	*tinydom.Element
	ulElement *tinydom.Element
}

// New creates a new instance of Nav
func New() *Nav {
	doc := tinydom.GetDocument()
	element := doc.CreateElement("nav")
	ulElement := doc.CreateElement("ul")

	element.AppendChild(ulElement)

	return &Nav{
		Element:   element,
		ulElement: ulElement,
	}
}

func (n *Nav) AppendListItem(item *li.Li) *Nav {
	n.ulElement.AppendChild(item.Element)
	return n
}
