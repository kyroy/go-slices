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

package float64s_test

import (
	"github.com/kyroy/go-slices/float64s"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		s []float64
		f func(s float64) float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "basic",
			args: args{
				s: []float64{1, 2, -1, -3},
				f: func(s float64) float64 {
					return s + 1
				},
			},
			want: []float64{2, 3, 0, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []float64 = float64s.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
