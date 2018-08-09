// Copyright 2018 Dennis Kuhnert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package internal

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"Title": strings.Title,
}

// TypeTemplate is the template to generate {type}/type.go
var TypeTemplate = template.Must(template.New("type.go").Funcs(funcMap).Parse(`// Code generated by go generate; DO NOT EDIT.

// Copyright 2018 Dennis Kuhnert
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package {{ $.Package }} provides typical slice functions for the type {{ $.Type }}.
package {{ $.Package }}
{{ with $slice := $.Package | Title }}
// {{ $slice }} wraps []{{ $.Type }}.
type {{ $slice }} []{{ $.Type }}

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []{{ $.Type }}) {{ $slice }} {
	return {{ $slice }}(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func Map(s []{{ $.Type }}, f func(s {{ $.Type }}) {{ $.Type }}) {{ $slice }} {
	m := {{ $slice }}(make([]{{ $.Type }}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func (s {{ $slice }}) Map(f func(s {{ $.Type }}) {{ $.Type }}) {{ $slice }} {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []{{ $.Type }}, f func(s {{ $.Type }}) bool) {{ $slice }} {
	m := {{ $slice }}(make([]{{ $.Type }}, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s {{ $slice }}) Filter(f func(s {{ $.Type }}) bool) {{ $slice }} {
	return Filter(s, f)
}

// Reduce applies the provided function agains an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []{{ $.Type }}, f func(sum, value {{ $.Type }}) {{ $.Type }}, neutral {{ $.Type }}) {{ $.Type }} {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function agains an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s {{ $slice }}) Reduce(f func(sum, value {{ $.Type }}) {{ $.Type }}, neutral {{ $.Type }}) {{ $.Type }} {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []{{ $.Type }}) {{ $slice }} {
	m := {{ $slice }}(make([]{{ $.Type }}, 0, len(s)))
	seen := make(map[{{ $.Type }}]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s {{ $slice }}) Unique() {{ $slice }} {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []{{ $.Type }}, more ...[]{{ $.Type }}) {{ $slice }} {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[{{ $.Type }}]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[{{ $.Type }}]int)
		for _, e := range x {
			if v, ok := hits[e]; ok {
				hits[e] = v + 1
			} else {
				hits[e] = 1
			}
		}
		for k, v := range hits {
			if c, ok := counts[k]; ok {
				if c.maxHits > v {
					c.maxHits = v
				}
				c.slices++
				counts[k] = c
			}
		}
	}
	var m {{ $slice }}
	for e, c := range counts {
		if c.slices >= len(more)+1 {
			for i := 0; i < c.maxHits; i++ {
				m = append(m, e)
			}
		}
	}
	return m
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func (s {{ $slice }}) Intersect(more ...[]{{ $.Type }}) {{ $slice }} {
	return Intersect(s, more...)
}

// Contains returns true if the element is found in the slice.
func Contains(s []{{ $.Type }}, x {{ $.Type }}) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Contains returns true if the element is found in the slice.
func (s {{ $slice }}) Contains(x {{ $.Type }}) bool {
	return Contains(s, x)
}

// IndexOf returns the position of the first occurrence of the specified value.
func IndexOf(s []{{ $.Type }}, x {{ $.Type }}) int {
	for i, e := range s {
		if e == x {
			return i
		}
	}
	return -1
}

// IndexOf returns the position of the first occurrence of the specified value.
func (s {{ $slice }}) IndexOf(x {{ $.Type }}) int {
	return IndexOf(s, x)
}

// Find returns the first element that passes the test implemented by the provided function.
func Find(s []{{ $.Type }}, f func({{ $.Type }}) bool) *{{ $.Type }} {
	for _, e := range s {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Find returns the first element that passes the test implemented by the provided function.
func (s {{ $slice }}) Find(f func({{ $.Type }}) bool) *{{ $.Type }} {
	return Find(s, f)
}
{{ end }}`))
