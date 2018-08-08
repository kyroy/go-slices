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
	fmt.Println(y)
	fmt.Println(strings.Reduce(x, func(s, v string) string {
		return s + v
	}, ""))
	fmt.Println(strings.Unique(x))
}
