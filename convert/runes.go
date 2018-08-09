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

package convert

import (
	"github.com/kyroy/go-slices/bools"
	"github.com/kyroy/go-slices/float32s"
	"github.com/kyroy/go-slices/float64s"
	"github.com/kyroy/go-slices/interfaces"
	"github.com/kyroy/go-slices/int8s"
	"github.com/kyroy/go-slices/int16s"
	"github.com/kyroy/go-slices/int32s"
	"github.com/kyroy/go-slices/int64s"
	"github.com/kyroy/go-slices/uint8s"
	"github.com/kyroy/go-slices/uint16s"
	"github.com/kyroy/go-slices/uint32s"
	"github.com/kyroy/go-slices/uint64s"
	"github.com/kyroy/go-slices/ints"
	"github.com/kyroy/go-slices/uints"
	"github.com/kyroy/go-slices/bytes"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/strings"
)

// RunesBools creates a new slice with the results of calling the provided function on every element in the given array.
func RunesBools(s []rune, f func(s rune) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesFloat32s(s []rune, f func(s rune) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesFloat64s(s []rune, f func(s rune) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInterfaces(s []rune, f func(s rune) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInt8s(s []rune, f func(s rune) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInt16s(s []rune, f func(s rune) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInt32s(s []rune, f func(s rune) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInt64s(s []rune, f func(s rune) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUint8s(s []rune, f func(s rune) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUint16s(s []rune, f func(s rune) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUint32s(s []rune, f func(s rune) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUint64s(s []rune, f func(s rune) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesInts creates a new slice with the results of calling the provided function on every element in the given array.
func RunesInts(s []rune, f func(s rune) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUints creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUints(s []rune, f func(s rune) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesBytes creates a new slice with the results of calling the provided function on every element in the given array.
func RunesBytes(s []rune, f func(s rune) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesComplex64s(s []rune, f func(s rune) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func RunesComplex128s(s []rune, f func(s rune) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func RunesUintptrs(s []rune, f func(s rune) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// RunesStrings creates a new slice with the results of calling the provided function on every element in the given array.
func RunesStrings(s []rune, f func(s rune) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}
