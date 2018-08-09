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
// limitations under the License.package strslice_test

package main

import (
	"fmt"
	"github.com/kyroy/go-slices/strings"
)

func main() {
	x := []string{"a", "b", "a", "c", "d"}
	y := strings.Filter(x, func(s string) bool {
		return s > "b"
	}).Map(func(s string) string {
		return s + "!"
	})
	z := strings.Unique(x)
	fmt.Println("x           ", x)
	fmt.Println("y filter map", y)
	fmt.Println("z unique    ", z)

	a := strings.New([]string{"e", "f", "g", "h"})
	b := a.Filter(func(s string) bool {
		return s > "b"
	}).Map(func(s string) string {
		return s + "!"
	})
	c := a.Reduce(func(s, v string) string {
		return s + v
	}, "")
	fmt.Println("a new       ", a)
	fmt.Println("b filter map", b)
	fmt.Println("c reduce    ", c)
}
