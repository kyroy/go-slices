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

package int16s_test

import (
	"fmt"
	"github.com/kyroy/go-slices/int16s"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = fmt.Print

func TestMap(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 3, 4, 6},
				f: func(x int16) int16 { return x + 1 },
			},
			want: []int16{2, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = int16s.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = int16s.New(tt.args.s).Map(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) bool
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3, 4, 2},
				f: func(x int16) bool { return x > 1 && x < 4 },
			},
			want: []int16{2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = int16s.Filter(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = int16s.New(tt.args.s).Filter(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		s       []int16
		f       func(sum, value int16) int16
		neutral int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "basic",
			args: args{
				s:       []int16{1, 3, 4},
				f:       func(sum, value int16) int16 { return sum + value },
				neutral: 0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int16 = int16s.Reduce(tt.args.s, tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
			got = int16s.New(tt.args.s).Reduce(tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		s []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 3, 4, 6, 3, 2, 1},
			},
			want: []int16{1, 3, 4, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = int16s.Unique(tt.args.s)
			assert.Equal(t, tt.want, got)
			got = int16s.New(tt.args.s).Unique()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		s    []int16
		more [][]int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "same",
			args: args{
				s: []int16{1, 2, 3, 4},
				more: [][]int16{{1, 2, 3, 4}},
			},
			want: []int16{1, 2, 3, 4},
		},{
			name: "multiple",
			args: args{
				s: []int16{1, 1, 2, 1, 3, 5, 4},
				more: [][]int16{{1, 2, 2, 3, 1}, {1, 2, 1, 5, 4}},
			},
			want: []int16{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = int16s.Intersect(tt.args.s, tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
			got = int16s.New(tt.args.s).Intersect(tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
