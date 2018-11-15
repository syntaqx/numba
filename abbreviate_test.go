package numba

import (
	"testing"
)

type abbreviateTest struct {
	in  int64
	out string
}

var abbreviateTests = []abbreviateTest{
	{-1, "-1"},
	{0, "0"},
	{1, "1"},
	{100, "100"},
	{1000, "1k"},
	{1100, "1.1k"},
	{1111, "1.1k"},
	{1500, "1.5k"},
	{1000000, "1M"},
	{1500000, "1.5M"},
	{1000000000, "1B"},
	{1500000000, "1.5B"},
	{1000000000000, "1T"},
	{1500000000000, "1.5T"},
	{1000000000000000, "1q"},
	{1000000000000000000, "1Q"},
}

func TestAbbreviate(t *testing.T) {
	t.Parallel()
	for _, tt := range abbreviateTests {
		func(tt abbreviateTest) {
			t.Run(tt.out, func(t *testing.T) {
				t.Parallel()
				s := Abbreviate(tt.in)
				if s != tt.out {
					t.Errorf("got %q, want %q from %d", s, tt.out, tt.in)
				}
			})
		}(tt)
	}
}
