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

package float64s_test

import (
	"fmt"
	"github.com/kyroy/go-slices/float64s"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = fmt.Print

func TestMap(t *testing.T) {
	type args struct {
		s []float64
		f func(s float64) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []float64{1, 3, 4, 6},
				f: func(x float64) float64 { return x + 1 },
			},
			want: []float64{2, 4, 5, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = float64s.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).Map(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		s []float64
		f func(s float64) bool
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []float64{1, 2, 3, 4, 2},
				f: func(x float64) bool { return x > 1 && x < 4 },
			},
			want: []float64{2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = float64s.Filter(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).Filter(tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestReduce(t *testing.T) {
	type args struct {
		s       []float64
		f       func(sum, value float64) float64
		neutral float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic",
			args: args{
				s:       []float64{1, 3, 4},
				f:       func(sum, value float64) float64 { return sum + value },
				neutral: 0,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got float64 = float64s.Reduce(tt.args.s, tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).Reduce(tt.args.f, tt.args.neutral)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	type args struct {
		s []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []float64{1, 3, 4, 6, 3, 2, 1},
			},
			want: []float64{1, 3, 4, 6, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = float64s.Unique(tt.args.s)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).Unique()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		s    []float64
		more [][]float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "same",
			args: args{
				s:    []float64{1, 2, 3, 4},
				more: [][]float64{{1, 2, 3, 4}},
			},
			want: []float64{1, 2, 3, 4},
		}, {
			name: "multiple",
			args: args{
				s:    []float64{1, 1, 2, 1, 3, 5, 4},
				more: [][]float64{{1, 2, 2, 3, 1}, {1, 2, 1, 5, 4}},
			},
			want: []float64{1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = float64s.Intersect(tt.args.s, tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
			got = float64s.New(tt.args.s).Intersect(tt.args.more...)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestContains(t *testing.T) {
	type args struct {
		s []float64
		x float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contains",
			args: args{
				s: []float64{1, 2, 3, 4},
				x: 3,
			},
			want: true,
		}, {
			name: "not_contains",
			args: args{
				s: []float64{1, 1, 2, 1, 3, 5, 4},
				x: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got bool = float64s.Contains(tt.args.s, tt.args.x)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).Contains(tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIndexOf(t *testing.T) {
	type args struct {
		s []float64
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contains",
			args: args{
				s: []float64{1, 2, 3, 4},
				x: 3,
			},
			want: 2,
		}, {
			name: "not_contains",
			args: args{
				s: []float64{1, 1, 2, 1, 3, 5, 4},
				x: 0,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got int = float64s.IndexOf(tt.args.s, tt.args.x)
			assert.Equal(t, tt.want, got)
			got = float64s.New(tt.args.s).IndexOf(tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFind(t *testing.T) {
	type args struct {
		s []float64
		f func(float64) bool
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		found bool
	}{
		{
			name: "contains",
			args: args{
				s: []float64{1, 2, 3, 4, 2},
				f: func(x float64) bool { return x > 1 && x < 4 },
			},
			want:  2,
			found: true,
		}, {
			name: "not_contains",
			args: args{
				s: []float64{1, 2, 3, 4, 2},
				f: func(x float64) bool { return x > 5 },
			},
			want:  0,
			found: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got1 *float64 = float64s.Find(tt.args.s, tt.args.f)
			var got2 *float64 = float64s.New(tt.args.s).Find(tt.args.f)
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
