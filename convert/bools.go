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

// Package convert provides type safe map functions between different types.
package convert

import (
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
	"github.com/kyroy/go-slices/runes"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/strings"
)

// BoolsFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsFloat32s(s []bool, f func(s bool) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsFloat64s(s []bool, f func(s bool) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInterfaces(s []bool, f func(s bool) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt8s(s []bool, f func(s bool) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt16s(s []bool, f func(s bool) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt32s(s []bool, f func(s bool) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt64s(s []bool, f func(s bool) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint8s(s []bool, f func(s bool) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint16s(s []bool, f func(s bool) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint32s(s []bool, f func(s bool) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint64s(s []bool, f func(s bool) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInts creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInts(s []bool, f func(s bool) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUints creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUints(s []bool, f func(s bool) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsBytes creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsBytes(s []bool, f func(s bool) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsRunes creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsRunes(s []bool, f func(s bool) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsComplex64s(s []bool, f func(s bool) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsComplex128s(s []bool, f func(s bool) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUintptrs(s []bool, f func(s bool) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsStrings creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsStrings(s []bool, f func(s bool) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}
