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
	"github.com/kyroy/go-slices/uints"
	"github.com/kyroy/go-slices/bytes"
	"github.com/kyroy/go-slices/runes"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/strings"
)

// IntsBools creates a new slice with the results of calling the provided function on every element in the given array.
func IntsBools(s []int, f func(s int) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsFloat32s(s []int, f func(s int) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsFloat64s(s []int, f func(s int) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInterfaces(s []int, f func(s int) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt8s(s []int, f func(s int) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt16s(s []int, f func(s int) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt32s(s []int, f func(s int) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt64s(s []int, f func(s int) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint8s(s []int, f func(s int) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint16s(s []int, f func(s int) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint32s(s []int, f func(s int) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint64s(s []int, f func(s int) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUints creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUints(s []int, f func(s int) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsBytes creates a new slice with the results of calling the provided function on every element in the given array.
func IntsBytes(s []int, f func(s int) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsRunes creates a new slice with the results of calling the provided function on every element in the given array.
func IntsRunes(s []int, f func(s int) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsComplex64s(s []int, f func(s int) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsComplex128s(s []int, f func(s int) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUintptrs(s []int, f func(s int) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsStrings creates a new slice with the results of calling the provided function on every element in the given array.
func IntsStrings(s []int, f func(s int) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}
