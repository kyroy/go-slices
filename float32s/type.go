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

// Package float32s provides typical slice functions for the type float32.
package float32s

// Float32s wraps []float32.
type Float32s []float32

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []float32) Float32s {
	return Float32s(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func Map(s []float32, f func(s float32) float32) Float32s {
	m := Float32s(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func (s Float32s) Map(f func(s float32) float32) Float32s {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []float32, f func(s float32) bool) Float32s {
	m := Float32s(make([]float32, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s Float32s) Filter(f func(s float32) bool) Float32s {
	return Filter(s, f)
}

// Reduce applies the provided function against an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []float32, f func(sum, value float32) float32, neutral float32) float32 {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s Float32s) Reduce(f func(sum, value float32) float32, neutral float32) float32 {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []float32) Float32s {
	m := Float32s(make([]float32, 0, len(s)))
	seen := make(map[float32]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s Float32s) Unique() Float32s {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []float32, more ...[]float32) Float32s {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[float32]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[float32]int)
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
	var m Float32s
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
func (s Float32s) Intersect(more ...[]float32) Float32s {
	return Intersect(s, more...)
}

// Contains returns true if the element is found in the slice.
func Contains(s []float32, x float32) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Contains returns true if the element is found in the slice.
func (s Float32s) Contains(x float32) bool {
	return Contains(s, x)
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func IndexOf(s []float32, x float32) int {
	for i, e := range s {
		if e == x {
			return i
		}
	}
	return -1
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func (s Float32s) IndexOf(x float32) int {
	return IndexOf(s, x)
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func Find(s []float32, f func(float32) bool) *float32 {
	for _, e := range s {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func (s Float32s) Find(f func(float32) bool) *float32 {
	return Find(s, f)
}
