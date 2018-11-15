package numba

import (
	"testing"
)

var bytesTests = []struct {
	Value     int64
	Base      int
	Precision int
	Result    string
}{
	{1000, 10, 0, "1KB"},
	{1024, 10, 2, "1.02KB"},
	{1024, 2, 1, "1.0KiB"},
	{56000, 10, 2, "56.00KB"},        // 1 sec of 56K modem
	{1024000, 2, 2, "1000.00KiB"},    // A "1MB" floppy disk is 1000KiB
	{52428800, 2, 0, "50MiB"},        // A "50MB" file on a CD is 50MiB
	{1300000000, 2, 2, "1.21GiB"},    // A 1.3GB file on a DVD is 1.3GB
	{300000000000, 10, 1, "300.0GB"}, // A 300GB hard disk is 300GB
	{8589934592, 2, 2, "8.00GiB"},    // "8GB" of RAM is 8GiB
}

func TestBytes(t *testing.T) {
	t.Parallel()

	for _, test := range bytesTests {
		tt := test
		t.Run(tt.Result, func(t *testing.T) {
			t.Parallel()
			ts := Bytes(tt.Value, tt.Base, tt.Precision)
			if ts != tt.Result {
				t.Errorf("%d base %d precision %d should be %s, was %s", tt.Value,
					tt.Base, tt.Precision, tt.Result, ts)
			}
		})
	}
}

var parseBytesTests = []struct {
	Value  string
	Result int64
}{
	{"1KB", 1000},
	{"1KiB", 1024},
	{"6GiB", 6 * 1024 * 1024 * 1024},
}

func TestParseBytes(t *testing.T) {
	t.Parallel()

	for _, test := range parseBytesTests {
		tt := test
		t.Run(tt.Value, func(t *testing.T) {
			t.Parallel()
			x, err := ParseBytes(tt.Value)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if x != tt.Result {
				t.Errorf("failed to parse %s correctly", tt.Value)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	t.Parallel()

	t.Run("10MB", func(t *testing.T) {
		t.Parallel()
		x, y := split("10MB")
		if x != "10" || y != "MB" {
			t.Errorf("failed to split 10MB")
		}
	})

	t.Run("5 GiB", func(t *testing.T) {
		t.Parallel()
		x, y := split("5 GiB")
		if x != "5" || y != "GiB" {
			t.Errorf("failed to split 5 GiB")
		}
	})
}

func TestErrors(t *testing.T) {
	t.Parallel()

	t.Run("1ZB", func(t *testing.T) {
		_, err := ParseBytes("1ZB")
		if err == nil {
			t.Errorf("failed to flag range error")
		}
	})
}
