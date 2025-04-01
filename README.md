# iecbyte [![Go Report Card](https://goreportcard.com/badge/github.com/andrewheberle/iecbyte?logo=go&style=flat-square)](https://goreportcard.com/report/github.com/andrewheberle/iecbyte) [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/andrewheberle/iecbyte)

This package provides a `Flag` type that can be used as a custom flag for `flag` and `github.com/sp13/pflag` as it satisifes the `flag.Value` and `pflag.Value` interfaces.

Command line flag values may be as either a bare number such as 1024, or using IEC byte suffixes such as Ki, Mi, Gi etc.

Using this the following values are all equivalent:
* 2097152
* 2048Ki
* 2Mi

Numbers must be specified as non-negative integers and the supported (case sensitive) suffixes are:
* Ki (1024 bytes)
* Mi (1024 * 1024 bytes)
* Gi (1024 * 1024 * 1024 bytes)
* Ti (1024 * 1024 * 1024 * 1024 bytes)
* Pi (1024 * 1024 * 1024 * 1024 * 1024 bytes)
* Ei (1024 * 1024 * 1024 * 1024 * 1024 * 1024 bytes)

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
