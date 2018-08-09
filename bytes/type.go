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

// Package bytes provides typical slice functions for the type byte.
package bytes

// Bytes wraps []byte.
type Bytes []byte

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []byte) Bytes {
	return Bytes(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the calling array.
func Map(s []byte, f func(s byte) byte) Bytes {
	m := Bytes(make([]byte, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the calling array.
func (s Bytes) Map(f func(s byte) byte) Bytes {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []byte, f func(s byte) bool) Bytes {
	m := Bytes(make([]byte, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s Bytes) Filter(f func(s byte) bool) Bytes {
	return Filter(s, f)
}

// Reduce applies the provided function agains an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []byte, f func(sum, value byte) byte, neutral byte) byte {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function agains an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s Bytes) Reduce(f func(sum, value byte) byte, neutral byte) byte {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []byte) Bytes {
	m := Bytes(make([]byte, 0, len(s)))
	seen := make(map[byte]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s Bytes) Unique() Bytes {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []byte, more ...[]byte) Bytes {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[byte]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[byte]int)
		for _, e := range x {
			if v, ok := hits[e]; ok {
				hits[e] = v + 1
			} else {
				hits[e] = 1
			}
		}
		for k, v := range hits {
			if c, ok := counts[k]; ok {
				if c.maxHits > v {
					c.maxHits = v
				}
				c.slices++
				counts[k] = c
			}
		}
	}
	var m Bytes
	for e, c := range counts {
		if c.slices >= len(more)+1 {
			for i := 0; i < c.maxHits; i++ {
				m = append(m, e)
			}
		}
	}
	return m
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func (s Bytes) Intersect(more ...[]byte) Bytes {
	return Intersect(s, more...)
}
