package numba

import (
	"math"
	"strconv"
)

var abbreviations = []string{"k", "M", "B", "T", "q", "Q", "s", "S"}

// Abbreviate formats integer n to it's large scale abbreviated format.
func Abbreviate(n int) string {
	num := strconv.Itoa(n)
	l := len(num)

	if n > 0 && l > 3 {
		f := float64(n)
		f /= math.Pow(10, float64(l-1))
		f = math.Floor(f*10.0+0.5) / 10.0

		i := l/3 - 1
		if len(abbreviations) >= i {
			return strconv.FormatFloat(f, 'f', -1, 64) + abbreviations[i]
		}
	}

	return num
}
