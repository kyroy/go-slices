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
	"github.com/kyroy/go-slices/uint16s"
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

// Uint8sBools creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sBools(s []uint8, f func(s uint8) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sFloat32s(s []uint8, f func(s uint8) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sFloat64s(s []uint8, f func(s uint8) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInterfaces(s []uint8, f func(s uint8) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInt8s(s []uint8, f func(s uint8) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInt16s(s []uint8, f func(s uint8) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInt32s(s []uint8, f func(s uint8) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInt64s(s []uint8, f func(s uint8) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sUint16s(s []uint8, f func(s uint8) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sUint32s(s []uint8, f func(s uint8) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sUint64s(s []uint8, f func(s uint8) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sInts creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sInts(s []uint8, f func(s uint8) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sUints creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sUints(s []uint8, f func(s uint8) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sBytes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sBytes(s []uint8, f func(s uint8) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sRunes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sRunes(s []uint8, f func(s uint8) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sComplex64s(s []uint8, f func(s uint8) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sComplex128s(s []uint8, f func(s uint8) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sUintptrs(s []uint8, f func(s uint8) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint8sStrings creates a new slice with the results of calling the provided function on every element in the given array.
func Uint8sStrings(s []uint8, f func(s uint8) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}