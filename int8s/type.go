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

// Package int8s provides typical slice functions for the type int8.
package int8s

// Int8s wraps []int8.
type Int8s []int8

// New is a convenient function that wraps the given slice to be able to call the struct package functions directly.
func New(s []int8) Int8s {
	return Int8s(s)
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func Map(s []int8, f func(s int8) int8) Int8s {
	m := Int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map creates a new slice with the results of calling the provided function on every element in the given array.
func (s Int8s) Map(f func(s int8) int8) Int8s {
	return Map(s, f)
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func Filter(s []int8, f func(s int8) bool) Int8s {
	m := Int8s(make([]int8, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter creates a new slice with all elements that pass the test implemented by the provided function.
func (s Int8s) Filter(f func(s int8) bool) Int8s {
	return Filter(s, f)
}

// Reduce applies the provided function against an accumulator and each element in the array (from left to right) to reduce it to a single value.
func Reduce(s []int8, f func(sum, value int8) int8, neutral int8) int8 {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce applies the provided function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func (s Int8s) Reduce(f func(sum, value int8) int8, neutral int8) int8 {
	return Reduce(s, f, neutral)
}

// Unique creates a new slice without duplicate values.
func Unique(s []int8) Int8s {
	m := Int8s(make([]int8, 0, len(s)))
	seen := make(map[int8]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique creates a new slice without duplicate values.
func (s Int8s) Unique() Int8s {
	return Unique(s)
}

// Intersect creates a new slice of values that are included in all given arrays.
// The order of values in the result is random.
func Intersect(s []int8, more ...[]int8) Int8s {
	type count struct {
		slices  int
		maxHits int
	}
	counts := make(map[int8]count)
	for _, e := range s {
		if c, ok := counts[e]; ok {
			c.maxHits++
			counts[e] = c
		} else {
			counts[e] = count{slices: 1, maxHits: 1}
		}
	}
	for _, x := range more {
		hits := make(map[int8]int)
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
	var m Int8s
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
func (s Int8s) Intersect(more ...[]int8) Int8s {
	return Intersect(s, more...)
}

// Contains returns true if the element is found in the slice.
func Contains(s []int8, x int8) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}

// Contains returns true if the element is found in the slice.
func (s Int8s) Contains(x int8) bool {
	return Contains(s, x)
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func IndexOf(s []int8, x int8) int {
	for i, e := range s {
		if e == x {
			return i
		}
	}
	return -1
}

// IndexOf returns the position of the first occurrence of the specified value.
// Returns -1 if not found.
func (s Int8s) IndexOf(x int8) int {
	return IndexOf(s, x)
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func Find(s []int8, f func(int8) bool) *int8 {
	for _, e := range s {
		if f(e) {
			return &e
		}
	}
	return nil
}

// Find returns the first element that passes the test implemented by the provided function.
// Returns nil if not found.
func (s Int8s) Find(f func(int8) bool) *int8 {
	return Find(s, f)
}
