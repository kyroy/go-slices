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
	"github.com/kyroy/go-slices/uint32s"
	"github.com/kyroy/go-slices/uint64s"
	"github.com/kyroy/go-slices/ints"
	"github.com/kyroy/go-slices/uints"
	"github.com/kyroy/go-slices/bytes"
	"github.com/kyroy/go-slices/runes"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/strings"
)

// Uint16sBools creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sBools(s []uint16, f func(s uint16) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sFloat32s(s []uint16, f func(s uint16) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sFloat64s(s []uint16, f func(s uint16) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInterfaces(s []uint16, f func(s uint16) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInt8s(s []uint16, f func(s uint16) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInt16s(s []uint16, f func(s uint16) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInt32s(s []uint16, f func(s uint16) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInt64s(s []uint16, f func(s uint16) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sUint8s(s []uint16, f func(s uint16) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sUint32s(s []uint16, f func(s uint16) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sUint64s(s []uint16, f func(s uint16) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sInts creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sInts(s []uint16, f func(s uint16) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sUints creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sUints(s []uint16, f func(s uint16) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sBytes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sBytes(s []uint16, f func(s uint16) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sRunes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sRunes(s []uint16, f func(s uint16) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sComplex64s(s []uint16, f func(s uint16) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sComplex128s(s []uint16, f func(s uint16) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sUintptrs(s []uint16, f func(s uint16) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint16sStrings creates a new slice with the results of calling the provided function on every element in the given array.
func Uint16sStrings(s []uint16, f func(s uint16) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}
