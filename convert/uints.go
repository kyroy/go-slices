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
	"github.com/kyroy/go-slices/uint32s"
	"github.com/kyroy/go-slices/uint64s"
	"github.com/kyroy/go-slices/uint8s"
	"github.com/kyroy/go-slices/uintptrs"
)

// UintsBools creates a new slice with the results of calling the provided function on every element in the given array.
func UintsBools(s []uint, f func(s uint) bool) bools.Bools {
	m := bools.Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsBoolsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsBoolsE(s []uint, f func(s uint) (bool, error)) (bools.Bools, error) {
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

// UintsBoolsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsBoolsF(s []uint, f func(s uint) (bool, error)) bools.Bools {
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

// UintsFloat32s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsFloat32s(s []uint, f func(s uint) float32) float32s.Float32s {
	m := float32s.Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsFloat32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsFloat32sE(s []uint, f func(s uint) (float32, error)) (float32s.Float32s, error) {
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

// UintsFloat32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsFloat32sF(s []uint, f func(s uint) (float32, error)) float32s.Float32s {
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

// UintsFloat64s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsFloat64s(s []uint, f func(s uint) float64) float64s.Float64s {
	m := float64s.Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsFloat64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsFloat64sE(s []uint, f func(s uint) (float64, error)) (float64s.Float64s, error) {
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

// UintsFloat64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsFloat64sF(s []uint, f func(s uint) (float64, error)) float64s.Float64s {
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

// UintsInterfaces creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInterfaces(s []uint, f func(s uint) interface{}) interfaces.Interfaces {
	m := interfaces.Interfaces(make([]interface{}, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsInterfacesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsInterfacesE(s []uint, f func(s uint) (interface{}, error)) (interfaces.Interfaces, error) {
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

// UintsInterfacesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsInterfacesF(s []uint, f func(s uint) (interface{}, error)) interfaces.Interfaces {
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

// UintsInt8s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInt8s(s []uint, f func(s uint) int8) int8s.Int8s {
	m := int8s.Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsInt8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsInt8sE(s []uint, f func(s uint) (int8, error)) (int8s.Int8s, error) {
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

// UintsInt8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsInt8sF(s []uint, f func(s uint) (int8, error)) int8s.Int8s {
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

// UintsInt16s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInt16s(s []uint, f func(s uint) int16) int16s.Int16s {
	m := int16s.Int16s(make([]int16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsInt16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsInt16sE(s []uint, f func(s uint) (int16, error)) (int16s.Int16s, error) {
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

// UintsInt16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsInt16sF(s []uint, f func(s uint) (int16, error)) int16s.Int16s {
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

// UintsInt32s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInt32s(s []uint, f func(s uint) int32) int32s.Int32s {
	m := int32s.Int32s(make([]int32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsInt32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsInt32sE(s []uint, f func(s uint) (int32, error)) (int32s.Int32s, error) {
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

// UintsInt32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsInt32sF(s []uint, f func(s uint) (int32, error)) int32s.Int32s {
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

// UintsInt64s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInt64s(s []uint, f func(s uint) int64) int64s.Int64s {
	m := int64s.Int64s(make([]int64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsInt64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsInt64sE(s []uint, f func(s uint) (int64, error)) (int64s.Int64s, error) {
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

// UintsInt64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsInt64sF(s []uint, f func(s uint) (int64, error)) int64s.Int64s {
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

// UintsUint8s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsUint8s(s []uint, f func(s uint) uint8) uint8s.Uint8s {
	m := uint8s.Uint8s(make([]uint8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsUint8sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsUint8sE(s []uint, f func(s uint) (uint8, error)) (uint8s.Uint8s, error) {
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

// UintsUint8sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsUint8sF(s []uint, f func(s uint) (uint8, error)) uint8s.Uint8s {
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

// UintsUint16s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsUint16s(s []uint, f func(s uint) uint16) uint16s.Uint16s {
	m := uint16s.Uint16s(make([]uint16, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsUint16sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsUint16sE(s []uint, f func(s uint) (uint16, error)) (uint16s.Uint16s, error) {
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

// UintsUint16sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsUint16sF(s []uint, f func(s uint) (uint16, error)) uint16s.Uint16s {
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

// UintsUint32s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsUint32s(s []uint, f func(s uint) uint32) uint32s.Uint32s {
	m := uint32s.Uint32s(make([]uint32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsUint32sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsUint32sE(s []uint, f func(s uint) (uint32, error)) (uint32s.Uint32s, error) {
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

// UintsUint32sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsUint32sF(s []uint, f func(s uint) (uint32, error)) uint32s.Uint32s {
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

// UintsUint64s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsUint64s(s []uint, f func(s uint) uint64) uint64s.Uint64s {
	m := uint64s.Uint64s(make([]uint64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsUint64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsUint64sE(s []uint, f func(s uint) (uint64, error)) (uint64s.Uint64s, error) {
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

// UintsUint64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsUint64sF(s []uint, f func(s uint) (uint64, error)) uint64s.Uint64s {
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

// UintsInts creates a new slice with the results of calling the provided function on every element in the given array.
func UintsInts(s []uint, f func(s uint) int) ints.Ints {
	m := ints.Ints(make([]int, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsIntsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsIntsE(s []uint, f func(s uint) (int, error)) (ints.Ints, error) {
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

// UintsIntsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsIntsF(s []uint, f func(s uint) (int, error)) ints.Ints {
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

// UintsBytes creates a new slice with the results of calling the provided function on every element in the given array.
func UintsBytes(s []uint, f func(s uint) byte) bytes.Bytes {
	m := bytes.Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsBytesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsBytesE(s []uint, f func(s uint) (byte, error)) (bytes.Bytes, error) {
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

// UintsBytesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsBytesF(s []uint, f func(s uint) (byte, error)) bytes.Bytes {
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

// UintsRunes creates a new slice with the results of calling the provided function on every element in the given array.
func UintsRunes(s []uint, f func(s uint) rune) runes.Runes {
	m := runes.Runes(make([]rune, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsRunesE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsRunesE(s []uint, f func(s uint) (rune, error)) (runes.Runes, error) {
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

// UintsRunesF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsRunesF(s []uint, f func(s uint) (rune, error)) runes.Runes {
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

// UintsComplex64s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsComplex64s(s []uint, f func(s uint) complex64) complex64s.Complex64s {
	m := complex64s.Complex64s(make([]complex64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsComplex64sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsComplex64sE(s []uint, f func(s uint) (complex64, error)) (complex64s.Complex64s, error) {
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

// UintsComplex64sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsComplex64sF(s []uint, f func(s uint) (complex64, error)) complex64s.Complex64s {
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

// UintsComplex128s creates a new slice with the results of calling the provided function on every element in the given array.
func UintsComplex128s(s []uint, f func(s uint) complex128) complex128s.Complex128s {
	m := complex128s.Complex128s(make([]complex128, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsComplex128sE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsComplex128sE(s []uint, f func(s uint) (complex128, error)) (complex128s.Complex128s, error) {
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

// UintsComplex128sF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsComplex128sF(s []uint, f func(s uint) (complex128, error)) complex128s.Complex128s {
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

// UintsUintptrs creates a new slice with the results of calling the provided function on every element in the given array.
func UintsUintptrs(s []uint, f func(s uint) uintptr) uintptrs.Uintptrs {
	m := uintptrs.Uintptrs(make([]uintptr, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsUintptrsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsUintptrsE(s []uint, f func(s uint) (uintptr, error)) (uintptrs.Uintptrs, error) {
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

// UintsUintptrsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsUintptrsF(s []uint, f func(s uint) (uintptr, error)) uintptrs.Uintptrs {
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

// UintsStrings creates a new slice with the results of calling the provided function on every element in the given array.
func UintsStrings(s []uint, f func(s uint) string) strings.Strings {
	m := strings.Strings(make([]string, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// UintsStringsE creates a new slice with the results of calling the provided function on every element in the given array.
// Returns an error if one of the function calls fails.
func UintsStringsE(s []uint, f func(s uint) (string, error)) (strings.Strings, error) {
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

// UintsStringsF creates a new slice with the results of calling the provided function on every element in the given array.
// If the given function returns an error, the element will be ignored.
func UintsStringsF(s []uint, f func(s uint) (string, error)) strings.Strings {
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
