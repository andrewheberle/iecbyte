package iecbyte

import (
	"fmt"
	"testing"
)

type test struct {
	name      string
	f         *Flag
	value     string
	want      string
	wantInt64 uint64
	wantErr   bool
}

func tests() []test {
	return []test{
		// valid
		{"0", &Flag{0}, "0", "0", 0, false},
		{"1", &Flag{1}, "1", "1", 1, false},
		{"1023", &Flag{(1 * Kilobyte) - (1 * Byte)}, "1023", "1023", (1 * Kilobyte) - (1 * Byte), false},
		{"1024", &Flag{1 * Kilobyte}, "1024", "1Ki", 1 * Kilobyte, false},
		{"1Ki", &Flag{1 * Kilobyte}, "1Ki", "1Ki", 1 * Kilobyte, false},
		{"1025", &Flag{(1 * Kilobyte) + (1 * Byte)}, "1025", "1025", (1 * Kilobyte) + (1 * Byte), false},
		{"2048", &Flag{2 * Kilobyte}, "2048", "2Ki", 2 * Kilobyte, false},
		{"2Ki", &Flag{2 * Kilobyte}, "2Ki", "2Ki", 2 * Kilobyte, false},
		{"1048576", &Flag{1 * Megabyte}, "1048576", "1Mi", 1 * Megabyte, false},
		{"1Mi", &Flag{1 * Megabyte}, "1Mi", "1Mi", 1 * Megabyte, false},
		{"1049600", &Flag{(1 * Megabyte) + (1 * Kilobyte)}, "1049600", "1025Ki", (1 * Megabyte) + (1 * Kilobyte), false},
		{"1025Ki", &Flag{(1 * Megabyte) + (1 * Kilobyte)}, "1025Ki", "1025Ki", (1 * Megabyte) + (1 * Kilobyte), false},
		{"1050623", &Flag{1050623}, "1050623", "1050623", 1050623, false},
		{"1050624", &Flag{(1 * Megabyte) + (2 * Kilobyte)}, "1050624", "1026Ki", (1 * Megabyte) + (2 * Kilobyte), false},
		{"1026Ki", &Flag{(1 * Megabyte) + (2 * Kilobyte)}, "1026Ki", "1026Ki", (1 * Megabyte) + (2 * Kilobyte), false},
		{"2097152", &Flag{2 * Megabyte}, "2097152", "2Mi", 2 * Megabyte, false},
		{"2048Ki", &Flag{2 * Megabyte}, "2048Ki", "2Mi", 2 * Megabyte, false},
		{"2Mi", &Flag{2 * Megabyte}, "2Mi", "2Mi", 2 * Megabyte, false},
		{"1073741824", &Flag{1073741824}, "1073741824", "1Gi", 1073741824, false},
		{"1Gi", &Flag{1 * Gigabyte}, "1Gi", "1Gi", 1 * Gigabyte, false},
		{"1073741825", &Flag{1073741825}, "1073741825", "1073741825", 1073741825, false},
		{"1073742848", &Flag{1073742848}, "1073742848", "1048577Ki", 1073742848, false},
		{"1048577Ki", &Flag{1073742848}, "1048577Ki", "1048577Ki", 1073742848, false},
		{"2147483648", &Flag{2 * Gigabyte}, "2147483648", "2Gi", 2 * Gigabyte, false},
		{"2097152Ki", &Flag{2 * Gigabyte}, "2097152Ki", "2Gi", 2 * Gigabyte, false},
		{"2048Mi", &Flag{2 * Gigabyte}, "2048Mi", "2Gi", 2 * Gigabyte, false},
		{"2Gi", &Flag{2 * Gigabyte}, "2Gi", "2Gi", 2 * Gigabyte, false},
		{"10240Pi", &Flag{10240 * Petabyte}, "10Ei", "10Ei", 10240 * Petabyte, false},
		{"10Ei", &Flag{10 * Exabyte}, "10Ei", "10Ei", 10 * Exabyte, false},
		{"max", &Flag{18446744073709551615}, "18446744073709551615", "18446744073709551615", 18446744073709551615, false},

		// invalid
		{"-1", &Flag{}, "-1", "", 0, true},
		{"-1Ki", &Flag{}, "-1Ki", "", 0, true},
		{"1a", &Flag{}, "1a", "", 0, true},
		{"a", &Flag{}, "a", "", 0, true},
		{"1024mi", &Flag{}, "1024mi", "", 0, true},
		{"", &Flag{}, "", "", 0, true},
		{"Mi", &Flag{}, "Mi", "", 0, true},
	}
}

func TestFlag_Set(t *testing.T) {
	fmt.Printf("%s\n", &Flag{1073742848})
	for _, tt := range tests() {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.f.Set(tt.value); (err != nil) != tt.wantErr {
				t.Errorf("Test: %s, Flag.Set() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}

func TestFlag_String(t *testing.T) {
	for _, tt := range tests() {
		if tt.wantErr {
			continue
		}

		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.String(); got != tt.want {
				t.Errorf("Test: %s, Flag.String() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFlag_Get(t *testing.T) {
	for _, tt := range tests() {
		if tt.wantErr {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Get(); got != tt.wantInt64 {
				t.Errorf("Test: %s, Flag.Get() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestFlag_Misc(t *testing.T) {
	f := NewFlag(100)
	t.Run("Test NewFlag()", func(t *testing.T) {
		if f.String() != "100" {
			t.Errorf("NewFlag().String() = %v, want %v", f.String(), "100")
		}
	})
	t.Run("Test Get()", func(t *testing.T) {
		if f.Get() != 100 {
			t.Errorf("NewFlag().Get() = %v, want %v", f.Get(), 100)
		}
	})
	t.Run("Test Type()", func(t *testing.T) {
		if f.Type() != flagType {
			t.Errorf("NewFlag().Type() = %v, want %v", f.Type(), flagType)
		}
	})
}
