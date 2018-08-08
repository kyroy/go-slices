package internal

import (
	"text/template"
)

var TypeTemplate = template.Must(template.New("type.go").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at {{ .Timestamp.Format "2006 Jan 02 15:04:05 UTC" }}

// Package {{ .Package }} provides ...
package {{ .Package }}

// {{ .SliceType }} ...
type {{ .SliceType }} []{{ .Type }}

// Map ...
func Map(s []{{ .Type }}, f func(s {{ .Type }}) {{ .Type }}) {{ .SliceType }} {
	m := {{ .SliceType }}(make([]{{ .Type }}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map ...
func (s {{ .SliceType }}) Map(f func(s {{ .Type }}) {{ .Type }}) {{ .SliceType }} {
	return Map(s, f)
}

// Filter ...
func Filter(s []{{ .Type }}, f func(s {{ .Type }}) bool) {{ .SliceType }} {
	m := {{ .SliceType }}(make([]{{ .Type }}, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter ...
func (s {{ .SliceType }}) Filter(f func(s {{ .Type }}) bool) {{ .SliceType }} {
	return Filter(s, f)
}

// Reduce ...
func Reduce(s []{{ .Type }}, f func(sum, value {{ .Type }}) {{ .Type }}, neutral {{ .Type }}) {{ .Type }} {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce ...
func (s {{ .SliceType }}) Reduce(f func(sum, value {{ .Type }}) {{ .Type }}, neutral {{ .Type }}) {{ .Type }} {
	return Reduce(s, f, neutral)
}

// Unique ...
func Unique(s []{{ .Type }}) {{ .SliceType }} {
	m := {{ .SliceType }}(make([]{{ .Type }}, 0, len(s)))
	seen := make(map[{{ .Type }}]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique ...
func (s {{ .SliceType }}) Unique() {{ .SliceType }} {
	return Unique(s)
}
`))
