package iecbyte_test

import (
	"flag"
	"fmt"

	"github.com/andrewheberle/iecbyte"
)

func ExampleNewFlag() {
	size := iecbyte.NewFlag(1024 * 1024)

	flag.Var(&size, "size", "Size in IEC bytes")
	//
	// In this example flag.Parse() is commented out as this forms part of the tests
	// of this module, so parsing the command line flags is disabled.
	//
	// In a real program you would need to call flag.Parse()
	//
	// flag.Parse()

	fmt.Printf("Size is %s\n", size)
	// Output: Size is 1Mi
}
