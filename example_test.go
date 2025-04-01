package iecbyte_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/andrewheberle/iecbyte"
	"github.com/spf13/pflag"
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

func ExampleFlag_Type() {
	// This example doesn't explicitly show the calling of Type(),
	// however the pflag package calls Type() as part of it's display
	// of command line flag defaults that is being manually called below.

	size := iecbyte.NewFlag(1024 * 1024)

	fs := pflag.NewFlagSet("example", pflag.ExitOnError)
	fs.SetOutput(os.Stdout)
	fs.Var(&size, "size", "Size in IEC bytes")
	fs.Parse([]string{})
	fs.PrintDefaults()
	// Output:
	//		--size bytes (IEC)   Size in IEC bytes (default 1Mi)
}
