# go-slices

[![GoDoc](https://godoc.org/github.com/kyroy/go-slices?status.svg)](https://godoc.org/github.com/kyroy/go-slices)
[![Build Status](https://travis-ci.org/kyroy/go-slices.svg?branch=master)](https://travis-ci.org/kyroy/go-slices)
[![Codecov](https://img.shields.io/codecov/c/github/kyroy/go-slices.svg)](https://codecov.io/gh/kyroy/go-slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyroy/go-slices)](https://goreportcard.com/report/github.com/kyroy/go-slices)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kyroy/go-slices/blob/master/LICENSE)

A type safe utility library for Go slices, providing contains, find, map, filter, reduce, unique, intersect, indexOf.

```bash
go get github.com/kyroy/go-slices/...
```

## Example using `strings`

```go
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
```
prints
```
x            [a b a c d]
y filter map [c! d!]
z unique     [a b c d]
a new        [e f g h]
b filter map [e! f! g! h!]
c reduce     efgh
```

## Map between different types with `convert` or `interfaces`

```go
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
```
prints
```
a [1 2 3 4]
b [1? 2? 3? 4?]
c [1 4 6 8]
d [1 4 6 8]
e [1 4 6 8] err <nil>
f [1 a 2]
g [1 2]
h [] err strconv.Atoi: parsing "a": invalid syntax
x [1 2 3 4]
y [1! 2! 3! 4!]
```

## Generating Slices
```bash
go run internal/generator/main.go internal/generator/types
```
