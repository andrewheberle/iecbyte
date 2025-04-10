package iecbyte

import (
	"fmt"
	"strconv"
	"strings"
)

// Flag satisfies the flag.Value and pflag.Value interfaces
type Flag struct {
	n uint64
}

type multiplier struct {
	Suffix string
	Value  uint64
}

const (
	_ uint64 = iota
	Byte
	Kilobyte = uint64(1 << (10 * (iota - 1)))
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

var multipliers = []multiplier{
	{"Ei", Exabyte},
	{"Pi", Petabyte},
	{"Ti", Terabyte},
	{"Gi", Gigabyte},
	{"Mi", Megabyte},
	{"Ki", Kilobyte},
}

const flagType = "bytes (IEC)"

// NewFlag is used to initialise a new iecbyte.Flag with a default value
func NewFlag(n uint64) Flag {
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
func (f Flag) Get() uint64 {
	return f.n
}

func parse(v string, m uint64) (uint64, error) {
	// parse the provided value (v) as an uint64
	n, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, err
	}

	// return the value of the parsed int64 multiplied by the multiplier (m)
	return n * m, nil
}
