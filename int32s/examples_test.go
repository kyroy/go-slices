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

package int32s_test

import (
	"fmt"
	"github.com/kyroy/go-slices/int32s"
)

func ExampleFilter() {
	fmt.Println(int32s.Filter([]int32{1, 2, 3, 4}, func(x int32) bool {
      return x % 2 == 0
    }))
	// Output: [2 4]
}

func ExampleIntersect() {
	fmt.Println(int32s.Intersect([]int32{1, 2, 3, 3}, []int32{3, 2, 3}, []int32{3, 4, 3}))
	// Output: [3 3]
}

func ExampleMap() {
	fmt.Println(int32s.Map([]int32{1, 2, 3, 4}, func(x int32) int32 {
      return 2 * x
    }))
	// Output: [2 4 6 8]
}

func ExampleReduce() {
	fmt.Println(int32s.Reduce([]int32{1, 2, 3, 4}, func(x, y int32) int32 {
      return x + y
    }, 0))
	// Output: 10
}

func ExampleUnique() {
	fmt.Println(int32s.Unique([]int32{1, 2, 1, 3, 2, 4}))
	// Output: [1 2 3 4]
}
