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
	"github.com/kyroy/go-slices/convert"
	"github.com/kyroy/go-slices/interfaces"
	"strconv"
)

func main() {
	a := []int{1, 2, 3, 4}
	b := convert.IntsStrings(a, func(x int) string {
		return fmt.Sprintf("%d?", x)
	})
	fmt.Println("a", a)
	fmt.Println("b", b)

	c := []string{"1", "4", "6", "8"}
	d := convert.StringsIntsF(c, strconv.Atoi)
	e, err := convert.StringsIntsE(c, strconv.Atoi)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e, "err", err)

	f := []string{"1", "a", "2"}
	g := convert.StringsIntsF(f, strconv.Atoi)
	h, err := convert.StringsIntsE(f, strconv.Atoi)
	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h, "err", err)

	x := []interface{}{1, 2, 3, 4}
	y := interfaces.Map(x, func(s interface{}) interface{} {
		return fmt.Sprintf("%d!", s)
	})
	fmt.Println("x", x)
	fmt.Println("y", y)
}
