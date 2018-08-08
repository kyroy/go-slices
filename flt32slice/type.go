// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2018 Aug 08 11:02:02 UTC

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

// Package flt32slice provides ...
package flt32slice

// floats32 ...
type floats32 []float32

// Map ...
func Map(s []float32, f func(s float32) float32) floats32 {
	m := floats32(make([]float32, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map ...
func (s floats32) Map(f func(s float32) float32) floats32 {
	return Map(s, f)
}

// Filter ...
func Filter(s []float32, f func(s float32) bool) floats32 {
	m := floats32(make([]float32, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter ...
func (s floats32) Filter(f func(s float32) bool) floats32 {
	return Filter(s, f)
}

// Reduce ...
func Reduce(s []float32, f func(sum, value float32) float32, neutral float32) float32 {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce ...
func (s floats32) Reduce(f func(sum, value float32) float32, neutral float32) float32 {
	return Reduce(s, f, neutral)
}

// Unique ...
func Unique(s []float32) floats32 {
	m := floats32(make([]float32, 0, len(s)))
	seen := make(map[float32]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique ...
func (s floats32) Unique() floats32 {
	return Unique(s)
}
