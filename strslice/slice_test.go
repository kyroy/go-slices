package strslice_test

import (
	"github.com/kyroy/go-slices/strslice"
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
			var got []string = strslice.Map(tt.args.s, tt.args.f)
			assert.Equal(t, tt.want, got)
		})
	}
}
