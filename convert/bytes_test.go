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

func TestBytesBools(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) bool { return x > 1 },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.BytesBools(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesBoolsE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool
			got, _ = convert.BytesBoolsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesBoolsF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (bool, error)
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (bool, error) { return x > 1, nil },
			},
			want: []bool{false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = convert.BytesBoolsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat32s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) float32 { return float32(x) },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.BytesFloat32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat32sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32
			got, _ = convert.BytesFloat32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat32sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (float32, error)
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (float32, error) { return float32(x), nil },
			},
			want: []float32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float32 = convert.BytesFloat32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat64s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) float64 { return float64(x) },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.BytesFloat64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat64sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64
			got, _ = convert.BytesFloat64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesFloat64sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (float64, error)
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (float64, error) { return float64(x), nil },
			},
			want: []float64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = convert.BytesFloat64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInterfaces(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) interface{} { return int(x) },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.BytesInterfaces(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInterfacesE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{}
			got, _ = convert.BytesInterfacesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInterfacesF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (interface{}, error)
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (interface{}, error) { return int(x), nil },
			},
			want: []interface{}{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []interface{} = convert.BytesInterfacesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt8s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) int8 { return int8(x) },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.BytesInt8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt8sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8
			got, _ = convert.BytesInt8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt8sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int8, error)
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int8, error) { return int8(x), nil },
			},
			want: []int8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int8 = convert.BytesInt8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt16s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) int16 { return int16(x) },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = convert.BytesInt16s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt16sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int16, error)
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int16, error) { return int16(x), nil },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16
			got, _ = convert.BytesInt16sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt16sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int16, error)
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int16, error) { return int16(x), nil },
			},
			want: []int16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int16 = convert.BytesInt16sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt32s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) int32 { return int32(x) },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.BytesInt32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt32sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32
			got, _ = convert.BytesInt32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt32sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int32, error)
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int32, error) { return int32(x), nil },
			},
			want: []int32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int32 = convert.BytesInt32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt64s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) int64 { return int64(x) },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.BytesInt64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt64sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64
			got, _ = convert.BytesInt64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInt64sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int64, error)
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int64, error) { return int64(x), nil },
			},
			want: []int64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int64 = convert.BytesInt64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint8s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) uint8 { return uint8(x) },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.BytesUint8s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint8sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8
			got, _ = convert.BytesUint8sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint8sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint8, error)
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint8, error) { return uint8(x), nil },
			},
			want: []uint8{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint8 = convert.BytesUint8sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint16s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) uint16 { return uint16(x) },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.BytesUint16s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint16sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16
			got, _ = convert.BytesUint16sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint16sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint16, error)
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint16, error) { return uint16(x), nil },
			},
			want: []uint16{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint16 = convert.BytesUint16sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint32s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) uint32 { return uint32(x) },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.BytesUint32s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint32sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32
			got, _ = convert.BytesUint32sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint32sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint32, error)
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint32, error) { return uint32(x), nil },
			},
			want: []uint32{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint32 = convert.BytesUint32sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint64s(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) uint64 { return uint64(x) },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.BytesUint64s(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint64sE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64
			got, _ = convert.BytesUint64sE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUint64sF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint64, error)
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint64, error) { return uint64(x), nil },
			},
			want: []uint64{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint64 = convert.BytesUint64sF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesInts(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) int { return int(x) },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.BytesInts(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesIntsE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			got, _ = convert.BytesIntsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesIntsF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (int, error)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (int, error) { return int(x), nil },
			},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int = convert.BytesIntsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUints(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) uint { return uint(x) },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint = convert.BytesUints(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUintsE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint, error)
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint, error) { return uint(x), nil },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint
			got, _ = convert.BytesUintsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesUintsF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (uint, error)
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (uint, error) { return uint(x), nil },
			},
			want: []uint{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []uint = convert.BytesUintsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesRunes(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) rune { return rune(x) },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.BytesRunes(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesRunesE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune
			got, _ = convert.BytesRunesE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesRunesF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (rune, error)
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (rune, error) { return rune(x), nil },
			},
			want: []rune{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []rune = convert.BytesRunesF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesStrings(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) string { return fmt.Sprint(x) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.BytesStrings(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesStringsE(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			got, _ = convert.BytesStringsE(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBytesStringsF(t *testing.T) {
	type args struct {
		s []byte
		f func(s byte) (string, error)
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []byte{1, 2, 3},
				f: func(x byte) (string, error) { return fmt.Sprint(x), nil },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = convert.BytesStringsF(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
