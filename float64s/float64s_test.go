// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2018 Aug 08 16:59:04 UTC

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
	"github.com/kyroy/go-slices/float64s"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
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
				f: func(x float64) bool { return x <= 3.4 },
			},
			want: []float64{1, 2, 3, 2},
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
