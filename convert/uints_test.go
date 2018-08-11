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

func TestUintsBools(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) bool { return x > 1 },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.UintsBools(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsBoolsE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool
			got, _ = convert.UintsBoolsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsBoolsF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.UintsBoolsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat32s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) float32 { return float32(x) },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.UintsFloat32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat32sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32
			got, _ = convert.UintsFloat32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat32sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.UintsFloat32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat64s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) float64 { return float64(x) },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.UintsFloat64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat64sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64
			got, _ = convert.UintsFloat64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsFloat64sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.UintsFloat64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInterfaces(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) interface{} { return int(x) },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.UintsInterfaces(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInterfacesE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{}
			got, _ = convert.UintsInterfacesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInterfacesF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.UintsInterfacesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt8s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) int8 { return int8(x) },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.UintsInt8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt8sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8
			got, _ = convert.UintsInt8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt8sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.UintsInt8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt16s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) int16 { return int16(x) },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = convert.UintsInt16s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt16sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int16, error)
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int16, error) { return int16(x), nil },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16
			got, _ = convert.UintsInt16sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt16sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int16, error)
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int16, error) { return int16(x), nil },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = convert.UintsInt16sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt32s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) int32 { return int32(x) },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.UintsInt32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt32sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32
			got, _ = convert.UintsInt32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt32sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.UintsInt32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt64s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) int64 { return int64(x) },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.UintsInt64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt64sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64
			got, _ = convert.UintsInt64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInt64sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.UintsInt64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint8s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) uint8 { return uint8(x) },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.UintsUint8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint8sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8
			got, _ = convert.UintsUint8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint8sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.UintsUint8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint16s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) uint16 { return uint16(x) },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.UintsUint16s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint16sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16
			got, _ = convert.UintsUint16sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint16sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.UintsUint16sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint32s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) uint32 { return uint32(x) },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.UintsUint32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint32sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32
			got, _ = convert.UintsUint32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint32sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.UintsUint32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint64s(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) uint64 { return uint64(x) },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.UintsUint64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint64sE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64
			got, _ = convert.UintsUint64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsUint64sF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.UintsUint64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsInts(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) int { return int(x) },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.UintsInts(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsIntsE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			got, _ = convert.UintsIntsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsIntsF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.UintsIntsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsBytes(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) byte { return byte(x) },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte = convert.UintsBytes(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsBytesE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (byte, error)
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (byte, error) { return byte(x), nil },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte
			got, _ = convert.UintsBytesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsBytesF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (byte, error)
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (byte, error) { return byte(x), nil },
			},
			want: []byte{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []byte = convert.UintsBytesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsRunes(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) rune { return rune(x) },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.UintsRunes(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsRunesE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune
			got, _ = convert.UintsRunesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsRunesF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.UintsRunesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsStrings(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) string { return fmt.Sprint(x) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.UintsStrings(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsStringsE(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			got, _ = convert.UintsStringsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUintsStringsF(t *testing.T) {
	type args struct {
		s []uint
		f func(s uint) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []uint{1, 2, 3},
				f: func(x uint) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.UintsStringsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
