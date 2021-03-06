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
	"github.com/kyroy/go-slices/bytes"
	"github.com/kyroy/go-slices/complex128s"
	"github.com/kyroy/go-slices/complex64s"
	"github.com/kyroy/go-slices/float32s"
	"github.com/kyroy/go-slices/float64s"
	"github.com/kyroy/go-slices/int16s"
	"github.com/kyroy/go-slices/int32s"
	"github.com/kyroy/go-slices/int64s"
	"github.com/kyroy/go-slices/int8s"
	"github.com/kyroy/go-slices/interfaces"
	"github.com/kyroy/go-slices/ints"
	"github.com/kyroy/go-slices/runes"
	"github.com/kyroy/go-slices/strings"
	"github.com/kyroy/go-slices/uint16s"
	"github.com/kyroy/go-slices/uint64s"
	"github.com/kyroy/go-slices/uint8s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/uints"
)

// Uint32sBools creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sBools(s []uint32, f func(s uint32) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sBoolsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sBoolsE(s []uint32, f func(s uint32) (bool, error)) (bools.Bools, error) {
	m := bools.Bools(make([]bool, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sBoolsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sBoolsF(s []uint32, f func(s uint32) (bool, error)) bools.Bools {
	m := bools.Bools(make([]bool, 0, len(s)))
	var (
		x   bool
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sFloat32s(s []uint32, f func(s uint32) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sFloat32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sFloat32sE(s []uint32, f func(s uint32) (float32, error)) (float32s.Float32s, error) {
	m := float32s.Float32s(make([]float32, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sFloat32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sFloat32sF(s []uint32, f func(s uint32) (float32, error)) float32s.Float32s {
	m := float32s.Float32s(make([]float32, 0, len(s)))
	var (
		x   float32
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sFloat64s(s []uint32, f func(s uint32) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sFloat64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sFloat64sE(s []uint32, f func(s uint32) (float64, error)) (float64s.Float64s, error) {
	m := float64s.Float64s(make([]float64, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sFloat64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sFloat64sF(s []uint32, f func(s uint32) (float64, error)) float64s.Float64s {
	m := float64s.Float64s(make([]float64, 0, len(s)))
	var (
		x   float64
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInterfaces(s []uint32, f func(s uint32) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sInterfacesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sInterfacesE(s []uint32, f func(s uint32) (interface{}, error)) (interfaces.Interfaces, error) {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sInterfacesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sInterfacesF(s []uint32, f func(s uint32) (interface{}, error)) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, 0, len(s)))
	var (
		x   interface{}
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInt8s(s []uint32, f func(s uint32) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sInt8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sInt8sE(s []uint32, f func(s uint32) (int8, error)) (int8s.Int8s, error) {
	m := int8s.Int8s(make([]int8, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sInt8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sInt8sF(s []uint32, f func(s uint32) (int8, error)) int8s.Int8s {
	m := int8s.Int8s(make([]int8, 0, len(s)))
	var (
		x   int8
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInt16s(s []uint32, f func(s uint32) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sInt16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sInt16sE(s []uint32, f func(s uint32) (int16, error)) (int16s.Int16s, error) {
	m := int16s.Int16s(make([]int16, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sInt16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sInt16sF(s []uint32, f func(s uint32) (int16, error)) int16s.Int16s {
	m := int16s.Int16s(make([]int16, 0, len(s)))
	var (
		x   int16
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInt32s(s []uint32, f func(s uint32) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sInt32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sInt32sE(s []uint32, f func(s uint32) (int32, error)) (int32s.Int32s, error) {
	m := int32s.Int32s(make([]int32, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sInt32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sInt32sF(s []uint32, f func(s uint32) (int32, error)) int32s.Int32s {
	m := int32s.Int32s(make([]int32, 0, len(s)))
	var (
		x   int32
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInt64s(s []uint32, f func(s uint32) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sInt64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sInt64sE(s []uint32, f func(s uint32) (int64, error)) (int64s.Int64s, error) {
	m := int64s.Int64s(make([]int64, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sInt64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sInt64sF(s []uint32, f func(s uint32) (int64, error)) int64s.Int64s {
	m := int64s.Int64s(make([]int64, 0, len(s)))
	var (
		x   int64
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sUint8s(s []uint32, f func(s uint32) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sUint8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sUint8sE(s []uint32, f func(s uint32) (uint8, error)) (uint8s.Uint8s, error) {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sUint8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sUint8sF(s []uint32, f func(s uint32) (uint8, error)) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, 0, len(s)))
	var (
		x   uint8
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sUint16s(s []uint32, f func(s uint32) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sUint16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sUint16sE(s []uint32, f func(s uint32) (uint16, error)) (uint16s.Uint16s, error) {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sUint16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sUint16sF(s []uint32, f func(s uint32) (uint16, error)) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, 0, len(s)))
	var (
		x   uint16
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sUint64s(s []uint32, f func(s uint32) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sUint64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sUint64sE(s []uint32, f func(s uint32) (uint64, error)) (uint64s.Uint64s, error) {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sUint64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sUint64sF(s []uint32, f func(s uint32) (uint64, error)) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, 0, len(s)))
	var (
		x   uint64
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sInts creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sInts(s []uint32, f func(s uint32) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sIntsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sIntsE(s []uint32, f func(s uint32) (int, error)) (ints.Ints, error) {
	m := ints.Ints(make([]int, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sIntsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sIntsF(s []uint32, f func(s uint32) (int, error)) ints.Ints {
	m := ints.Ints(make([]int, 0, len(s)))
	var (
		x   int
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sUints creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sUints(s []uint32, f func(s uint32) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sUintsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sUintsE(s []uint32, f func(s uint32) (uint, error)) (uints.Uints, error) {
	m := uints.Uints(make([]uint, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sUintsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sUintsF(s []uint32, f func(s uint32) (uint, error)) uints.Uints {
	m := uints.Uints(make([]uint, 0, len(s)))
	var (
		x   uint
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sBytes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sBytes(s []uint32, f func(s uint32) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sBytesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sBytesE(s []uint32, f func(s uint32) (byte, error)) (bytes.Bytes, error) {
	m := bytes.Bytes(make([]byte, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sBytesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sBytesF(s []uint32, f func(s uint32) (byte, error)) bytes.Bytes {
	m := bytes.Bytes(make([]byte, 0, len(s)))
	var (
		x   byte
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sRunes creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sRunes(s []uint32, f func(s uint32) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sRunesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sRunesE(s []uint32, f func(s uint32) (rune, error)) (runes.Runes, error) {
	m := runes.Runes(make([]rune, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sRunesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sRunesF(s []uint32, f func(s uint32) (rune, error)) runes.Runes {
	m := runes.Runes(make([]rune, 0, len(s)))
	var (
		x   rune
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sComplex64s(s []uint32, f func(s uint32) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sComplex64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sComplex64sE(s []uint32, f func(s uint32) (complex64, error)) (complex64s.Complex64s, error) {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sComplex64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sComplex64sF(s []uint32, f func(s uint32) (complex64, error)) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, 0, len(s)))
	var (
		x   complex64
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sComplex128s(s []uint32, f func(s uint32) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sComplex128sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sComplex128sE(s []uint32, f func(s uint32) (complex128, error)) (complex128s.Complex128s, error) {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sComplex128sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sComplex128sF(s []uint32, f func(s uint32) (complex128, error)) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, 0, len(s)))
	var (
		x   complex128
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sUintptrs(s []uint32, f func(s uint32) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sUintptrsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sUintptrsE(s []uint32, f func(s uint32) (uintptr, error)) (uintptrs.Uintptrs, error) {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sUintptrsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sUintptrsF(s []uint32, f func(s uint32) (uintptr, error)) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, 0, len(s)))
	var (
		x   uintptr
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}

// Uint32sStrings creates a new slice with the results of calling the provided function on every element in the given array.
func Uint32sStrings(s []uint32, f func(s uint32) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Uint32sStringsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func Uint32sStringsE(s []uint32, f func(s uint32) (string, error)) (strings.Strings, error) {
	m := strings.Strings(make([]string, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// Uint32sStringsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func Uint32sStringsF(s []uint32, f func(s uint32) (string, error)) strings.Strings {
	m := strings.Strings(make([]string, 0, len(s)))
	var (
		x   string
		err error
	)
	for _, v := range s {
		x, err = f(v)
		if err != nil {
			continue
		}
		m = append(m, x)
	}
	return m
}
