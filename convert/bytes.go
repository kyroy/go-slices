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
	"github.com/kyroy/go-slices/runes"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/strings"
)

// BytesBools creates a new slice with the results of calling the provided function on every element in the given array.
func BytesBools(s []byte, f func(s byte) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesFloat32s(s []byte, f func(s byte) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesFloat64s(s []byte, f func(s byte) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInterfaces(s []byte, f func(s byte) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInt8s(s []byte, f func(s byte) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInt16s(s []byte, f func(s byte) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInt32s(s []byte, f func(s byte) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInt64s(s []byte, f func(s byte) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUint8s(s []byte, f func(s byte) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUint16s(s []byte, f func(s byte) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUint32s(s []byte, f func(s byte) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUint64s(s []byte, f func(s byte) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesInts creates a new slice with the results of calling the provided function on every element in the given array.
func BytesInts(s []byte, f func(s byte) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUints creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUints(s []byte, f func(s byte) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesRunes creates a new slice with the results of calling the provided function on every element in the given array.
func BytesRunes(s []byte, f func(s byte) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesComplex64s(s []byte, f func(s byte) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func BytesComplex128s(s []byte, f func(s byte) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func BytesUintptrs(s []byte, f func(s byte) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BytesStrings creates a new slice with the results of calling the provided function on every element in the given array.
func BytesStrings(s []byte, f func(s byte) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}
