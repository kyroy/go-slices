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

package strings_test

import (
	"github.com/kyroy/go-slices/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		s []string
		f func(s string) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				s: []string{"a", "aa", "b", "c", "d"},
				f: func(s string) string {
					return s + "!"
				},
			},
			want: []string{"a!", "aa!", "b!", "c!", "d!"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string = strings.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
