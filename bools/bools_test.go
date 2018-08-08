package bools_test

import (
	"github.com/kyroy/go-slices/bools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		s []bool
		f func(s bool) bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				s: []bool{true, false, true},
				f: func(s bool) bool {
					return !s
				},
			},
			want: []bool{false, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []bool = bools.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
