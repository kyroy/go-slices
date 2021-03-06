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
	"github.com/kyroy/go-slices/uint32s"
	"github.com/kyroy/go-slices/uint64s"
	"github.com/kyroy/go-slices/uint8s"
	"github.com/kyroy/go-slices/uintptrs"
	"github.com/kyroy/go-slices/uints"
)

// BoolsFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsFloat32s(s []bool, f func(s bool) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsFloat32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsFloat32sE(s []bool, f func(s bool) (float32, error)) (float32s.Float32s, error) {
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

// BoolsFloat32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsFloat32sF(s []bool, f func(s bool) (float32, error)) float32s.Float32s {
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

// BoolsFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsFloat64s(s []bool, f func(s bool) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsFloat64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsFloat64sE(s []bool, f func(s bool) (float64, error)) (float64s.Float64s, error) {
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

// BoolsFloat64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsFloat64sF(s []bool, f func(s bool) (float64, error)) float64s.Float64s {
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

// BoolsInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInterfaces(s []bool, f func(s bool) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInterfacesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsInterfacesE(s []bool, f func(s bool) (interface{}, error)) (interfaces.Interfaces, error) {
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

// BoolsInterfacesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsInterfacesF(s []bool, f func(s bool) (interface{}, error)) interfaces.Interfaces {
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

// BoolsInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt8s(s []bool, f func(s bool) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsInt8sE(s []bool, f func(s bool) (int8, error)) (int8s.Int8s, error) {
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

// BoolsInt8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsInt8sF(s []bool, f func(s bool) (int8, error)) int8s.Int8s {
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

// BoolsInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt16s(s []bool, f func(s bool) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsInt16sE(s []bool, f func(s bool) (int16, error)) (int16s.Int16s, error) {
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

// BoolsInt16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsInt16sF(s []bool, f func(s bool) (int16, error)) int16s.Int16s {
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

// BoolsInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt32s(s []bool, f func(s bool) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsInt32sE(s []bool, f func(s bool) (int32, error)) (int32s.Int32s, error) {
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

// BoolsInt32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsInt32sF(s []bool, f func(s bool) (int32, error)) int32s.Int32s {
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

// BoolsInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInt64s(s []bool, f func(s bool) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsInt64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsInt64sE(s []bool, f func(s bool) (int64, error)) (int64s.Int64s, error) {
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

// BoolsInt64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsInt64sF(s []bool, f func(s bool) (int64, error)) int64s.Int64s {
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

// BoolsUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint8s(s []bool, f func(s bool) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUint8sE(s []bool, f func(s bool) (uint8, error)) (uint8s.Uint8s, error) {
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

// BoolsUint8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUint8sF(s []bool, f func(s bool) (uint8, error)) uint8s.Uint8s {
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

// BoolsUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint16s(s []bool, f func(s bool) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUint16sE(s []bool, f func(s bool) (uint16, error)) (uint16s.Uint16s, error) {
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

// BoolsUint16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUint16sF(s []bool, f func(s bool) (uint16, error)) uint16s.Uint16s {
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

// BoolsUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint32s(s []bool, f func(s bool) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUint32sE(s []bool, f func(s bool) (uint32, error)) (uint32s.Uint32s, error) {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	var err error
	for i, v := range s {
		m[i], err = f(v)
		if err != nil {
			return nil, err
		}
	}
	return m, nil
}

// BoolsUint32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUint32sF(s []bool, f func(s bool) (uint32, error)) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, 0, len(s)))
	var (
		x   uint32
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

// BoolsUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUint64s(s []bool, f func(s bool) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUint64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUint64sE(s []bool, f func(s bool) (uint64, error)) (uint64s.Uint64s, error) {
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

// BoolsUint64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUint64sF(s []bool, f func(s bool) (uint64, error)) uint64s.Uint64s {
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

// BoolsInts creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsInts(s []bool, f func(s bool) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsIntsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsIntsE(s []bool, f func(s bool) (int, error)) (ints.Ints, error) {
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

// BoolsIntsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsIntsF(s []bool, f func(s bool) (int, error)) ints.Ints {
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

// BoolsUints creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUints(s []bool, f func(s bool) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUintsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUintsE(s []bool, f func(s bool) (uint, error)) (uints.Uints, error) {
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

// BoolsUintsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUintsF(s []bool, f func(s bool) (uint, error)) uints.Uints {
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

// BoolsBytes creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsBytes(s []bool, f func(s bool) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsBytesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsBytesE(s []bool, f func(s bool) (byte, error)) (bytes.Bytes, error) {
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

// BoolsBytesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsBytesF(s []bool, f func(s bool) (byte, error)) bytes.Bytes {
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

// BoolsRunes creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsRunes(s []bool, f func(s bool) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsRunesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsRunesE(s []bool, f func(s bool) (rune, error)) (runes.Runes, error) {
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

// BoolsRunesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsRunesF(s []bool, f func(s bool) (rune, error)) runes.Runes {
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

// BoolsComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsComplex64s(s []bool, f func(s bool) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsComplex64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsComplex64sE(s []bool, f func(s bool) (complex64, error)) (complex64s.Complex64s, error) {
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

// BoolsComplex64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsComplex64sF(s []bool, f func(s bool) (complex64, error)) complex64s.Complex64s {
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

// BoolsComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsComplex128s(s []bool, f func(s bool) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsComplex128sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsComplex128sE(s []bool, f func(s bool) (complex128, error)) (complex128s.Complex128s, error) {
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

// BoolsComplex128sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsComplex128sF(s []bool, f func(s bool) (complex128, error)) complex128s.Complex128s {
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

// BoolsUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsUintptrs(s []bool, f func(s bool) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsUintptrsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsUintptrsE(s []bool, f func(s bool) (uintptr, error)) (uintptrs.Uintptrs, error) {
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

// BoolsUintptrsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsUintptrsF(s []bool, f func(s bool) (uintptr, error)) uintptrs.Uintptrs {
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

// BoolsStrings creates a new slice with the results of calling the provided function on every element in the given array.
func BoolsStrings(s []bool, f func(s bool) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// BoolsStringsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func BoolsStringsE(s []bool, f func(s bool) (string, error)) (strings.Strings, error) {
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

// BoolsStringsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func BoolsStringsF(s []bool, f func(s bool) (string, error)) strings.Strings {
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
