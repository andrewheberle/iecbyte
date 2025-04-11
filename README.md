# iecbyte

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewheberle/iecbyte?logo=go&style=flat-square)](https://goreportcard.com/report/github.com/andrewheberle/iecbyte)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/andrewheberle/iecbyte)
[![codecov](https://codecov.io/gh/andrewheberle/iecbyte/graph/badge.svg?token=MNFPOWU3VV)](https://codecov.io/gh/andrewheberle/iecbyte)

This package provides a `Flag` type that can be used as a custom flag for `flag` and `github.com/sp13/pflag` as it satisifes the `flag.Value` and `pflag.Value` interfaces.

Command line flag values may be as either a plain, non-negative integer or as a positive fixed-point number using one of these quantity suffixes: Ei, Pi, Ti, Gi, Mi, Ki.

For example the following all represent the same value:
* 2147483648
* 2097152Ki
* 2048Mi
* 2Gi

## Example

```go
package main

import (
    "flag"

    "github.com/andrewheberle/iecbyte"
)

func main() {
	size := iecbyte.NewFlag(1024 * 1024)

	flag.Var(&size, "size", "Size in IEC bytes")
	flag.Parse()

	fmt.Printf("Size is %s\n", size)
	// Output: Size is 1Mi
}
```
