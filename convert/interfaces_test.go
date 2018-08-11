// Code generated by go generate; DO NOT EDIT.

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

package convert_test

import (
	"fmt"
	"github.com/kyroy/go-slices/convert"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var _ = fmt.Print
var _ = strconv.Atoi

func TestInterfacesStrings(t *testing.T) {
	type args struct {
		s []interface{}
		f func(s interface{}) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []interface{}{1, 2, 3},
				f: func(x interface{}) string { return fmt.Sprint(x) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.InterfacesStrings(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterfacesStringsE(t *testing.T) {
	type args struct {
		s []interface{}
		f func(s interface{}) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []interface{}{1, 2, 3},
				f: func(x interface{}) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			got, _ = convert.InterfacesStringsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInterfacesStringsF(t *testing.T) {
	type args struct {
		s []interface{}
		f func(s interface{}) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []interface{}{1, 2, 3},
				f: func(x interface{}) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.InterfacesStringsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
