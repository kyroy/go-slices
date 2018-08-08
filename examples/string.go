package main

import (
	"fmt"
	"github.com/kyroy/go-slices/strslice"
)

func main() {
	x := []string{"a", "b", "a", "c", "d"}
	y := strslice.Filter(x, func(s string) bool {
		return s > "b"
	}).Map(func(s string) string {
		return s + "!"
	})
	fmt.Println(y)
	fmt.Println(strslice.Reduce(x, func(s, v string) string {
		return s + v
	}, ""))
	fmt.Println(strslice.Unique(x))
}
