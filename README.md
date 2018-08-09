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

	a := strings.New([]string{"e", "f", "g", "h"})
	b := a.Filter(func(s string) bool {
		return s > "b"
	}).Map(func(s string) string {
		return s + "!"
	})
	c := a.Reduce(func(s, v string) string {
		return s + v
	}, "")

	fmt.Println("x           ", x)
	fmt.Println("y filter map", y)
	fmt.Println("z unique    ", z)
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

## Generating Slices
```bash
go run generator/main.go generator/types
```
