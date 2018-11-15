package numba

import (
	"fmt"
	"testing"
)

var abbreviateTests = []struct {
	in  int
	out string
}{
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
	for _, tt := range abbreviateTests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			s := Abbreviate(tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q from %d", s, tt.out, tt.in)
			}
		})
	}
}
