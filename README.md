# go-slices

[![GoDoc](https://godoc.org/github.com/kyroy/go-slices?status.svg)](https://godoc.org/github.com/kyroy/go-slices)
[![Build Status](https://travis-ci.org/kyroy/go-slices.svg?branch=master)](https://travis-ci.org/kyroy/go-slices)
[![Codecov](https://img.shields.io/codecov/c/github/kyroy/go-slices.svg)](https://codecov.io/gh/kyroy/go-slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/kyroy/go-slices)](https://goreportcard.com/report/github.com/kyroy/go-slices)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/kyroy/go-slices/blob/master/LICENSE)

## Example using `strings`

```go
package main

import (
	"fmt"
	"github.com/kyroy/go-slices/strings"
)

func main() {
	x := []string{"a", "b", "a", "c", "d"}
	fmt.Println(strings.Filter(x, func(s string) bool {
		return s > "b"
	}).Map(func(s string) string {
		return s + "!"
	}))
	fmt.Println(strings.Reduce(x, func(s, v string) string {
		return s + v
	}, ""))
	fmt.Println(strings.Unique(x))
}
```
prints
```
[c! d!]
abacd
[a b c d]
```

## Generating Slices
```bash
go run generator/main.go TYPES.yaml
```
