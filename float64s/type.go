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

// Package float64s provides typical slice functions for the type float64.
package float64s

// Float64s wraps []float64.
type Float64s []float64

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []float64) Float64s {
	return Float64s(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func Map(s []float64, f func(s float64) float64) Float64s {
	m := Float64s(make([]float64, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func (s Float64s) Map(f func(s float64) float64) Float64s {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []float64, f func(s float64) bool) Float64s {
	m := Float64s(make([]float64, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s Float64s) Filter(f func(s float64) bool) Float64s {
	return Filter(s, f)
}

// Reduce applies the provided function against an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []float64, f func(sum, value float64) float64, neutral float64) float64 {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s Float64s) Reduce(f func(sum, value float64) float64, neutral float64) float64 {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []float64) Float64s {
	m := Float64s(make([]float64, 0, len(s)))
	seen := make(map[float64]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s Float64s) Unique() Float64s {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []float64, more ...[]float64) Float64s {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[float64]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[float64]int)
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
	var m Float64s
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
func (s Float64s) Intersect(more ...[]float64) Float64s {
	return Intersect(s, more...)
}

// Contains returns true if the element is found in the slice.
func Contains(s []float64, x float64) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Contains returns true if the element is found in the slice.
func (s Float64s) Contains(x float64) bool {
	return Contains(s, x)
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func IndexOf(s []float64, x float64) int {
	for i, e := range s {
		if e == x {
			return i
		}
	}
	return -1
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func (s Float64s) IndexOf(x float64) int {
	return IndexOf(s, x)
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func Find(s []float64, f func(float64) bool) *float64 {
	for _, e := range s {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func (s Float64s) Find(f func(float64) bool) *float64 {
	return Find(s, f)
}
