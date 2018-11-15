package numba

import (
	"fmt"
	"testing"
)

var abbreviateTests = []struct {
	in  int64
	out string
}{
	{-1, "-1"},
	{0, "0"},
	{1, "1"},
	{100, "100"},
	{1000, "1K"},
	{1100, "1.1K"},
	{1111, "1.1K"},
	{1500, "1.5K"},
	{1000000, "1M"},
	{1500000, "1.5M"},
	{1000000000, "1B"},
	{1500000000, "1.5B"},
	{1000000000000, "1T"},
	{1500000000000, "1.5T"},
}

func TestAbbreviate(t *testing.T) {
	for _, tt := range abbreviateTests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			s := Abbreviate(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q from %d", s, tt.out, tt.in)
			}
		})
	}
}
