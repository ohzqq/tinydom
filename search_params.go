package tinydom

import "syscall/js"

type SearchParams struct {
	js.Value
}

func NewSearchParams(search string) *SearchParams {
	return &SearchParams{Value: js.Global().Get("URLSearchParams").New(search)}
}

func (sp *SearchParams) Size() int {
	return sp.Call("size").Int()
}

func (sp *SearchParams) Get(key string) string {
	return sp.Call("get", key).String()
}

func (sp *SearchParams) Set(key, val string) {
	sp.Call("set", key, val)
}

func (sp *SearchParams) Append(key, val string) {
	sp.Call("append", key, val)
}

func (sp *SearchParams) Delete(key string, val ...string) {
	if len(val) > 0 {
		sp.Call("delete", key, val[0])
		return
	}
	sp.Call("delete", key)
}

func (sp *SearchParams) Has(key string) bool {
	return sp.Call("has", key).Bool()
}

func (sp *SearchParams) GetAll(key string) []string {
	all := sp.Call("getAll")
	vals := make([]string, all.Length())
	for i := 0; i < all.Length(); i++ {
		vals[i] = all.Index(i).String()
	}
	return vals
}

func (sp *SearchParams) Keys() []string {
	keys := sp.Call("keys")
	vals := make([]string, keys.Length())
	for i := 0; i < keys.Length(); i++ {
		vals[i] = keys.Index(i).String()
	}
	return vals
}

func (sp *SearchParams) Values() []string {
	values := sp.Call("values")
	vals := make([]string, values.Length())
	for i := 0; i < values.Length(); i++ {
		vals[i] = values.Index(i).String()
	}
	return vals
}
