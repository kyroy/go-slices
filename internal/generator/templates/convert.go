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
	"text/template"
)

// ConvertTemplate is the template to generate convert/{type}.go
var ConvertTemplate = template.Must(template.New("convert.go").Funcs(funcMap).Parse(`// Code generated by go generate; DO NOT EDIT.

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

package convert

import ({{ range .Other }}{{ if not (eq .Package $.Package) }}
	"github.com/kyroy/go-slices/{{ .Package }}"{{ end }}{{ end }}
)
{{ range .Other }}{{ if not (eq .Package $.Package) }}
// {{ $.Package | Title }}{{ .Package | Title }} creates a new slice with the results of calling the provided function on every element in the given array.
func {{ $.Package | Title }}{{ .Package | Title }}(s []{{ $.Type }}, f func(s {{ $.Type }}) {{ .Type }}) {{ .Package }}.{{ .Package | Title }} {
	m := {{ .Package }}.{{ .Package | Title }}(make([]{{ .Type }}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// {{ $.Package | Title }}{{ .Package | Title }}E creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func {{ $.Package | Title }}{{ .Package | Title }}E(s []{{ $.Type }}, f func(s {{ $.Type }}) ({{ .Type }}, error)) ({{ .Package }}.{{ .Package | Title }}, error) {
	m := {{ .Package }}.{{ .Package | Title }}(make([]{{ .Type }}, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// {{ $.Package | Title }}{{ .Package | Title }}F creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func {{ $.Package | Title }}{{ .Package | Title }}F(s []{{ $.Type }}, f func(s {{ $.Type }}) ({{ .Type }}, error)) {{ .Package }}.{{ .Package | Title }} {
	m := {{ .Package }}.{{ .Package | Title }}(make([]{{ .Type }}, 0, len(s)))
	var (
		x   {{ .Type }}
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}
{{ end }}{{ end }}`))
