package numba

import (
	"testing"
)

type bytesTest struct {
	value     int64
	base      int
	precision int
	out       string
}

var bytesTests = []bytesTest{
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
	for _, tt := range bytesTests {
		func(tt bytesTest) {
			t.Run(tt.out, func(t *testing.T) {
				t.Parallel()
				out := Bytes(tt.value, tt.base, tt.precision)
				if out != tt.out {
					t.Errorf("%d base %d precision 1%d should be %s, got %s", tt.value,
						tt.base, tt.precision, tt.out, out)
				}
			})
		}(tt)
	}
}

type parseBytesTest struct {
	in  string
	out int64
}

var parseBytesTests = []parseBytesTest{
	{"1KB", 1000},
	{"1KiB", 1024},
	{"6GiB", 6 * 1024 * 1024 * 1024},
}

func TestParseBytes(t *testing.T) {
	t.Parallel()
	for _, tt := range parseBytesTests {
		func(tt parseBytesTest) {
			t.Run(tt.in, func(t *testing.T) {
				t.Parallel()
				out, err := ParseBytes(tt.in)
				if err != nil {
					t.Errorf("error: %v", err)
				}
				if out != tt.out {
					t.Errorf("failed to parse %s correctly, expected %v got %v", tt.in, tt.out, out)
				}
			})
		}(tt)
	}
}

func TestParseBytesErrors(t *testing.T) {
	t.Parallel()
	_, err := ParseBytes("1ZB")
	if err == nil {
		t.Errorf("failed to flag range error")
	}
}

type byteSplitTest struct {
	in   string
	out1 string
	out2 string
}

var byteSplitTests = []byteSplitTest{
	{"10MB", "10", "MB"},
	{"5 GiB", "5", "GiB"},
}

func TestSplit(t *testing.T) {
	t.Parallel()
	for _, tt := range byteSplitTests {
		func(tt byteSplitTest) {
			t.Run(tt.in, func(t *testing.T) {
				t.Parallel()
				out1, out2 := byteSplit(tt.in)
				if out1 != tt.out1 || out2 != tt.out2 {
					t.Errorf("unexpected output from byteSplit. expected %s/%s, got %s/%s", tt.out1, tt.out2, out1, out2)
				}
			})
		}(tt)
	}
}
