# CountMin

[![Latest Version](http://img.shields.io/github/release/mtchavez/countmin.svg?style=flat-square)](https://github.com/mtchavez/countmin/releases)
[![Build Status](https://travis-ci.org/mtchavez/countmin.svg?branch=master)](https://travis-ci.org/mtchavez/countmin)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/countmin)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/countmin)](https://goreportcard.com/report/github.com/mtchavez/countmin)
[![Maintainability](https://api.codeclimate.com/v1/badges/9915ffa2990295885e85/maintainability)](https://codeclimate.com/github/mtchavez/countmin/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/9915ffa2990295885e85/test_coverage)](https://codeclimate.com/github/mtchavez/countmin/test_coverage)

CountMin sketching algorithm.

## Install

Install package

`go get -u github.com/mtchavez/countmin`

## Usage

```go
package main

import (
	"fmt"

	"github.com/mtchavez/countmin"
)

func main() {
	cm := countmin.New(10, 100000000)
	for _, i := range []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1} {
		cm.Add([]byte(fmt.Sprintf("%d", i)), int64(i))
	}
	fmt.Printf("Estimate of %d is %d\n", 1, cm.Count([]byte("1")))
	fmt.Printf("Estimate of %d is %d\n", 3, cm.Count([]byte("3")))
	fmt.Printf("Estimate of %d is %d\n", 9, cm.Count([]byte("9")))
	fmt.Println("Size: ", cm.Size())
	fmt.Println("Err: ", cm.RelativeError())
	fmt.Println("Confidence: ", cm.Confidence())
}
```

## Tests

Run using `go test ./... --cover`

Run benchmarks `go test --bench=.*`

## TODO

* Serialize/Deserialize
* TCP / HTTP server wrappers
