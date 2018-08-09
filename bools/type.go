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

// Package bools provides typical slice functions for the type bool.
package bools

// Bools wraps []bool.
type Bools []bool

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []bool) Bools {
	return Bools(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the calling array.
func Map(s []bool, f func(s bool) bool) Bools {
	m := Bools(make([]bool, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the calling array.
func (s Bools) Map(f func(s bool) bool) Bools {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []bool, f func(s bool) bool) Bools {
	m := Bools(make([]bool, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s Bools) Filter(f func(s bool) bool) Bools {
	return Filter(s, f)
}

// Reduce applies the provided function agains an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []bool, f func(sum, value bool) bool, neutral bool) bool {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function agains an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s Bools) Reduce(f func(sum, value bool) bool, neutral bool) bool {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []bool) Bools {
	m := Bools(make([]bool, 0, len(s)))
	seen := make(map[bool]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s Bools) Unique() Bools {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []bool, more ...[]bool) Bools {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[bool]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[bool]int)
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
	var m Bools
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
func (s Bools) Intersect(more ...[]bool) Bools {
	return Intersect(s, more...)
}

// Contains returns true if the element is found in the slice.
func Contains(s []bool, x bool) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Contains returns true if the element is found in the slice.
func (s Bools) Contains(x bool) bool {
	return Contains(s, x)
}

// IndexOf returns the position of the first occurrence of the specified value.
func IndexOf(s []bool, x bool) int {
	for i, e := range s {
		if e == x {
			return i
		}
	}
	return -1
}

// IndexOf returns the position of the first occurrence of the specified value.
func (s Bools) IndexOf(x bool) int {
	return IndexOf(s, x)
}

// Find returns the first element that passes the test implemented by the provided function.
func Find(s []bool, f func(bool) bool) *bool {
	for _, e := range s {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Find returns the first element that passes the test implemented by the provided function.
func (s Bools) Find(f func(bool) bool) *bool {
	return Find(s, f)
}
