// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2018 Aug 08 16:59:04 UTC

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

// Package int8s provides ...
package int8s

// int8s ...
type int8s []int8

// New ...
func New(s []int8) int8s {
	return int8s(s)
}

// Map ...
func Map(s []int8, f func(s int8) int8) int8s {
	m := int8s(make([]int8, len(s)))
	for i, v := range s {
		m[i] = f(v)
	}
	return m
}

// Map ...
func (s int8s) Map(f func(s int8) int8) int8s {
	return Map(s, f)
}

// Filter ...
func Filter(s []int8, f func(s int8) bool) int8s {
	m := int8s(make([]int8, 0, len(s)))
	for _, v := range s {
		if f(v) {
			m = append(m, v)
		}
	}
	return m
}

// Filter ...
func (s int8s) Filter(f func(s int8) bool) int8s {
	return Filter(s, f)
}

// Reduce ...
func Reduce(s []int8, f func(sum, value int8) int8, neutral int8) int8 {
	res := neutral
	for _, e := range s {
		res = f(res, e)
	}
	return res
}

// Reduce ...
func (s int8s) Reduce(f func(sum, value int8) int8, neutral int8) int8 {
	return Reduce(s, f, neutral)
}

// Unique ...
func Unique(s []int8) int8s {
	m := int8s(make([]int8, 0, len(s)))
	seen := make(map[int8]struct{})
	for _, v := range s {
		if _, ok := seen[v]; !ok {
			m = append(m, v)
			seen[v] = struct{}{}
		}
	}
	return m
}

// Unique ...
func (s int8s) Unique() int8s {
	return Unique(s)
}
