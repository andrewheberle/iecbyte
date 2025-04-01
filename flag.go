package iecbyte

import (
	"fmt"
	"strconv"
	"strings"
)

// Flag satisfies the pflag.Value interface
type Flag struct {
	n int64
}

type multiplier struct {
	Suffix string
	Value  int64
}

var multipliers = []multiplier{
	{"Ei", 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
	{"Pi", 1024 * 1024 * 1024 * 1024 * 1024},
	{"Ti", 1024 * 1024 * 1024 * 1024},
	{"Gi", 1024 * 1024 * 1024},
	{"Mi", 1024 * 1024},
	{"Ki", 1024},
}

const flagType = "bytes (IEC)"

// NewFlag is used to initialise a new iecbyte.Flag with a default value
//
// A value of n that is < 0 will be set to 0
func NewFlag(n int64) Flag {
	if n < 0 {
		n = 0
	}
	return Flag{n}
}

func (f Flag) String() string {
	for _, m := range multipliers {
		if f.n >= m.Value && f.n%m.Value == 0 {
			return fmt.Sprintf("%d%s", f.n/m.Value, m.Suffix)
		}
	}

	return fmt.Sprintf("%d", f.n)
}

func (f *Flag) Set(value string) error {
	// loop over multipliers and see if the value has the suffix in question
	for _, m := range multipliers {
		if strings.HasSuffix(value, m.Suffix) {
			// trim the suffix and parse the remainder of the string
			v := strings.TrimSuffix(value, m.Suffix)
			n, err := parse(v, m.Value)
			if err != nil {
				return err
			}

			// set the value of the flag
			f.n = n

			return nil
		}
	}

	// at this point the string either had no suffix on an unsupported one so parse as-is
	n, err := parse(value, 1)
	if err != nil {
		return err
	}

	// set the value of the flag
	f.n = n

	return nil
}

// Type returns a user facing type for this flag and is required to satisfy the pflag.Value interface
func (f Flag) Type() string {
	return flagType
}

// Get returns the value of the Flag as an int64
func (f Flag) Get() int64 {
	return f.n
}

func parse(v string, m int64) (int64, error) {
	// parse the provided value (v) as an int64
	n, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, err
	}

	// reject negative values
	if n < 0 {
		return 0, fmt.Errorf("cannot be negative")
	}

	// return the value of the parsed int64 multiplied by the multiplier (m)
	return n * m, nil
}
