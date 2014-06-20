# CountMin

[![Build Status](https://drone.io/github.com/mtchavez/countmin/status.png)](https://drone.io/github.com/mtchavez/countmin/latest)
[![Coverage Status](https://coveralls.io/repos/mtchavez/countmin/badge.png?branch=master)](https://coveralls.io/r/mtchavez/countmin?branch=master)

CountMin sketching algorithm.

## Install

Install package

`go get -u github.com/mtchavez/countmin/countmin`

## Usage

```go
package main

import (
	"fmt"

	"github.com/mtchavez/countmin/countmin"
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

## Docs

Docs on [GoDoc](http://godoc.org/github.com/mtchavez/countmin/countmin)

## Tests

Run using `go test ./... --cover`

## TODO

* Merge functionality
* Serialize/Deserialize
* Benchmarks
* TCP / HTTP server wrappers

## License

Written by Chavez

Released under the MIT License: http://www.opensource.org/licenses/mit-license.php
