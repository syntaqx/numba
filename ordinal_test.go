package numba

import (
	"testing"
)

type ordinalTest struct {
	in  int64
	out string
}

var ordinalTests = []ordinalTest{
	{0, "0th"},
	{1, "1st"},
	{2, "2nd"},
	{3, "3rd"},
	{4, "4th"},
	{10, "10th"},
	{11, "11th"},
	{12, "12th"},
	{13, "13th"},
	{101, "101st"},
	{102, "102nd"},
	{103, "103rd"},
}

func TestOrdinal(t *testing.T) {
	t.Parallel()
	for _, tt := range ordinalTests {
		func(tt ordinalTest) {
			t.Run(tt.out, func(t *testing.T) {
				t.Parallel()
				s := Ordinal(tt.in)
				if s != tt.out {
					t.Errorf("got %q, want %q from %d", s, tt.out, tt.in)
				}
			})
		}(tt)
	}
}
