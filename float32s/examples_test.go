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

package float32s_test

import (
	"fmt"
	"github.com/kyroy/go-slices/float32s"
)

func ExampleMap() {
	fmt.Println(float32s.Map([]float32{1, 2, 3}, func(x float32) float32 {
      return 2 * x
    }))
	// Output: [2 4 6]
}
