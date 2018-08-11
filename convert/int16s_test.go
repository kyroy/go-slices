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

func TestInt16sBools(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) bool { return x > 1 },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.Int16sBools(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sBoolsE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool
			got, _ = convert.Int16sBoolsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sBoolsF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.Int16sBoolsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat32s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) float32 { return float32(x) },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.Int16sFloat32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat32sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32
			got, _ = convert.Int16sFloat32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat32sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.Int16sFloat32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat64s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) float64 { return float64(x) },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.Int16sFloat64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat64sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64
			got, _ = convert.Int16sFloat64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sFloat64sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.Int16sFloat64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInterfaces(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) interface{} { return int(x) },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.Int16sInterfaces(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInterfacesE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{}
			got, _ = convert.Int16sInterfacesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInterfacesF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.Int16sInterfacesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt8s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) int8 { return int8(x) },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.Int16sInt8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt8sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8
			got, _ = convert.Int16sInt8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt8sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.Int16sInt8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt32s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) int32 { return int32(x) },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.Int16sInt32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt32sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32
			got, _ = convert.Int16sInt32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt32sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.Int16sInt32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt64s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) int64 { return int64(x) },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.Int16sInt64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt64sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64
			got, _ = convert.Int16sInt64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInt64sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.Int16sInt64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint8s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) uint8 { return uint8(x) },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.Int16sUint8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint8sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8
			got, _ = convert.Int16sUint8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint8sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.Int16sUint8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint16s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) uint16 { return uint16(x) },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.Int16sUint16s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint16sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16
			got, _ = convert.Int16sUint16sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint16sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.Int16sUint16sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint32s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) uint32 { return uint32(x) },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.Int16sUint32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint32sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32
			got, _ = convert.Int16sUint32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint32sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.Int16sUint32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint64s(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) uint64 { return uint64(x) },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.Int16sUint64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint64sE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64
			got, _ = convert.Int16sUint64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUint64sF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.Int16sUint64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sInts(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) int { return int(x) },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.Int16sInts(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sIntsE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			got, _ = convert.Int16sIntsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sIntsF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.Int16sIntsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUints(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) uint { return uint(x) },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint = convert.Int16sUints(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUintsE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint, error)
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint, error) { return uint(x), nil },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint
			got, _ = convert.Int16sUintsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sUintsF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (uint, error)
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (uint, error) { return uint(x), nil },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint = convert.Int16sUintsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sBytes(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) byte { return byte(x) },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte = convert.Int16sBytes(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sBytesE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (byte, error)
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (byte, error) { return byte(x), nil },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			got, _ = convert.Int16sBytesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sBytesF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (byte, error)
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (byte, error) { return byte(x), nil },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte = convert.Int16sBytesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sRunes(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) rune { return rune(x) },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.Int16sRunes(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sRunesE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune
			got, _ = convert.Int16sRunesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sRunesF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.Int16sRunesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sStrings(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) string { return fmt.Sprint(x) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.Int16sStrings(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sStringsE(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			got, _ = convert.Int16sStringsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestInt16sStringsF(t *testing.T) {
	type args struct {
		s []int16
		f func(s int16) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []int16{1, 2, 3},
				f: func(x int16) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.Int16sStringsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
