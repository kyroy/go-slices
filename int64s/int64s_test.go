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

package int64s_test

import (
	"fmt"
	"github.com/kyroy/go-slices/int64s"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = fmt.Print

func TestMap(t *testing.T) {
	type args struct {
		s []int64
		f func(s int64) int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int64{1, 3, 4, 6},
				f: func(x int64) int64 { return x + 1 },
			},
			want: []int64{2, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = int64s.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).Map(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		s []int64
		f func(s int64) bool
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int64{1, 2, 3, 4, 2},
				f: func(x int64) bool { return x > 1 && x < 4 },
			},
			want: []int64{2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = int64s.Filter(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).Filter(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		s       []int64
		f       func(sum, value int64) int64
		neutral int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "basic",
			args: args{
				s:       []int64{1, 3, 4},
				f:       func(sum, value int64) int64 { return sum + value },
				neutral: 0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int64 = int64s.Reduce(tt.args.s, tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).Reduce(tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		s []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int64{1, 3, 4, 6, 3, 2, 1},
			},
			want: []int64{1, 3, 4, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = int64s.Unique(tt.args.s)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).Unique()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		s    []int64
		more [][]int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "same",
			args: args{
				s:    []int64{1, 2, 3, 4},
				more: [][]int64{{1, 2, 3, 4}},
			},
			want: []int64{1, 2, 3, 4},
		}, {
			name: "multiple",
			args: args{
				s:    []int64{1, 1, 2, 1, 3, 5, 4},
				more: [][]int64{{1, 2, 2, 3, 1}, {1, 2, 1, 5, 4}},
			},
			want: []int64{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = int64s.Intersect(tt.args.s, tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
			got = int64s.New(tt.args.s).Intersect(tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains",
			args: args{
				s: []int64{1, 2, 3, 4},
				x: 3,
			},
			want: true,
		}, {
			name: "not_contains",
			args: args{
				s: []int64{1, 1, 2, 1, 3, 5, 4},
				x: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool = int64s.Contains(tt.args.s, tt.args.x)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).Contains(tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIndexOf(t *testing.T) {
	type args struct {
		s []int64
		x int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contains",
			args: args{
				s: []int64{1, 2, 3, 4},
				x: 3,
			},
			want: 2,
		}, {
			name: "not_contains",
			args: args{
				s: []int64{1, 1, 2, 1, 3, 5, 4},
				x: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int = int64s.IndexOf(tt.args.s, tt.args.x)
			assert.Equal(t, tt.want, got)
			got = int64s.New(tt.args.s).IndexOf(tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		s []int64
		f func(int64) bool
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		found bool
	}{
		{
			name: "contains",
			args: args{
				s: []int64{1, 2, 3, 4, 2},
				f: func(x int64) bool { return x > 1 && x < 4 },
			},
			want:  2,
			found: true,
		}, {
			name: "not_contains",
			args: args{
				s: []int64{1, 2, 3, 4, 2},
				f: func(x int64) bool { return x > 5 },
			},
			want:  0,
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got1 *int64 = int64s.Find(tt.args.s, tt.args.f)
			var got2 *int64 = int64s.New(tt.args.s).Find(tt.args.f)
			if tt.found {
				assert.Equal(t, &tt.want, got1)
				assert.Equal(t, &tt.want, got2)
			} else {
				assert.Nil(t, got1)
				assert.Nil(t, got2)
			}
		})
	}
}
