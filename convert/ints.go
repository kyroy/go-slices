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

// IntsBoolsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsBoolsE(s []int, f func(s int) (bool, error)) (bools.Bools, error) {
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

// IntsBoolsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsBoolsF(s []int, f func(s int) (bool, error)) bools.Bools {
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

// IntsFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsFloat32s(s []int, f func(s int) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsFloat32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsFloat32sE(s []int, f func(s int) (float32, error)) (float32s.Float32s, error) {
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

// IntsFloat32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsFloat32sF(s []int, f func(s int) (float32, error)) float32s.Float32s {
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

// IntsFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsFloat64s(s []int, f func(s int) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsFloat64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsFloat64sE(s []int, f func(s int) (float64, error)) (float64s.Float64s, error) {
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

// IntsFloat64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsFloat64sF(s []int, f func(s int) (float64, error)) float64s.Float64s {
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

// IntsInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInterfaces(s []int, f func(s int) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInterfacesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsInterfacesE(s []int, f func(s int) (interface{}, error)) (interfaces.Interfaces, error) {
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

// IntsInterfacesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsInterfacesF(s []int, f func(s int) (interface{}, error)) interfaces.Interfaces {
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

// IntsInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt8s(s []int, f func(s int) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsInt8sE(s []int, f func(s int) (int8, error)) (int8s.Int8s, error) {
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

// IntsInt8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsInt8sF(s []int, f func(s int) (int8, error)) int8s.Int8s {
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

// IntsInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt16s(s []int, f func(s int) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsInt16sE(s []int, f func(s int) (int16, error)) (int16s.Int16s, error) {
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

// IntsInt16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsInt16sF(s []int, f func(s int) (int16, error)) int16s.Int16s {
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

// IntsInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt32s(s []int, f func(s int) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsInt32sE(s []int, f func(s int) (int32, error)) (int32s.Int32s, error) {
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

// IntsInt32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsInt32sF(s []int, f func(s int) (int32, error)) int32s.Int32s {
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

// IntsInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsInt64s(s []int, f func(s int) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsInt64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsInt64sE(s []int, f func(s int) (int64, error)) (int64s.Int64s, error) {
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

// IntsInt64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsInt64sF(s []int, f func(s int) (int64, error)) int64s.Int64s {
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

// IntsUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint8s(s []int, f func(s int) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUint8sE(s []int, f func(s int) (uint8, error)) (uint8s.Uint8s, error) {
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

// IntsUint8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUint8sF(s []int, f func(s int) (uint8, error)) uint8s.Uint8s {
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

// IntsUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint16s(s []int, f func(s int) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUint16sE(s []int, f func(s int) (uint16, error)) (uint16s.Uint16s, error) {
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

// IntsUint16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUint16sF(s []int, f func(s int) (uint16, error)) uint16s.Uint16s {
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

// IntsUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint32s(s []int, f func(s int) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUint32sE(s []int, f func(s int) (uint32, error)) (uint32s.Uint32s, error) {
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

// IntsUint32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUint32sF(s []int, f func(s int) (uint32, error)) uint32s.Uint32s {
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

// IntsUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUint64s(s []int, f func(s int) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUint64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUint64sE(s []int, f func(s int) (uint64, error)) (uint64s.Uint64s, error) {
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

// IntsUint64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUint64sF(s []int, f func(s int) (uint64, error)) uint64s.Uint64s {
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

// IntsUints creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUints(s []int, f func(s int) uint) uints.Uints {
	m := uints.Uints(make([]uint, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUintsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUintsE(s []int, f func(s int) (uint, error)) (uints.Uints, error) {
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

// IntsUintsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUintsF(s []int, f func(s int) (uint, error)) uints.Uints {
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

// IntsBytes creates a new slice with the results of calling the provided function on every element in the given array.
func IntsBytes(s []int, f func(s int) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsBytesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsBytesE(s []int, f func(s int) (byte, error)) (bytes.Bytes, error) {
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

// IntsBytesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsBytesF(s []int, f func(s int) (byte, error)) bytes.Bytes {
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

// IntsRunes creates a new slice with the results of calling the provided function on every element in the given array.
func IntsRunes(s []int, f func(s int) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsRunesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsRunesE(s []int, f func(s int) (rune, error)) (runes.Runes, error) {
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

// IntsRunesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsRunesF(s []int, f func(s int) (rune, error)) runes.Runes {
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

// IntsComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsComplex64s(s []int, f func(s int) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsComplex64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsComplex64sE(s []int, f func(s int) (complex64, error)) (complex64s.Complex64s, error) {
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

// IntsComplex64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsComplex64sF(s []int, f func(s int) (complex64, error)) complex64s.Complex64s {
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

// IntsComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func IntsComplex128s(s []int, f func(s int) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsComplex128sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsComplex128sE(s []int, f func(s int) (complex128, error)) (complex128s.Complex128s, error) {
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

// IntsComplex128sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsComplex128sF(s []int, f func(s int) (complex128, error)) complex128s.Complex128s {
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

// IntsUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func IntsUintptrs(s []int, f func(s int) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsUintptrsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsUintptrsE(s []int, f func(s int) (uintptr, error)) (uintptrs.Uintptrs, error) {
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

// IntsUintptrsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsUintptrsF(s []int, f func(s int) (uintptr, error)) uintptrs.Uintptrs {
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

// IntsStrings creates a new slice with the results of calling the provided function on every element in the given array.
func IntsStrings(s []int, f func(s int) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// IntsStringsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func IntsStringsE(s []int, f func(s int) (string, error)) (strings.Strings, error) {
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

// IntsStringsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func IntsStringsF(s []int, f func(s int) (string, error)) strings.Strings {
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
