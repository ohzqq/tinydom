package tinydom

import "syscall/js"

type URL struct {
	js.Value
}

func ParseURL(args ...string) (*URL, bool) {
	var u, b string
	switch len(args) {
	case 0:
		return nil, false
	case 2:
		b = args[1]
		fallthrough
	case 1:
		u = args[0]
	}
	return &URL{
		Value: js.Global().Get("URL").Call("parse", u, b),
	}, true
}

func (u *URL) Hash() string {
	return u.Get("hash").String()
}

func (u *URL) Host() string {
	return u.Get("host").String()
}

func (u *URL) Hostname() string {
	return u.Get("hostname").String()
}

func (u *URL) Href() string {
	return u.Get("href").String()
}

func (u *URL) Origin() string {
	return u.Get("origin").String()
}

func (u *URL) Pathname() string {
	return u.Get("pathname").String()
}

func (u *URL) Port() string {
	return u.Get("port").String()
}

func (u *URL) Protocol() string {
	return u.Get("protocol").String()
}

func (u *URL) Search() string {
	return u.Get("search").String()
}

func (u *URL) SearchParams() *SearchParams {
	return &SearchParams{Value: u.Get("searchParams")}
}
