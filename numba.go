package numba

import (
	"math"
	"strconv"
)

var abbrSuffix = [4]string{"K", "M", "B", "T"}

// NumAbbreviate converts integer n to a friendly text abbreviation.
// For example, 1000 becomes '1K' and 1500 becomes '1.5k'.
// If n is less than 100, the number is just returned as a string.
// Accepts octal or hex numbers, returned value is decimal in those cases too.
func Abbreviate(n int64) string {
	num := strconv.FormatInt(n, 10)
	l := len(num)

	if n > 0 && l > 3 {
		f := float64(n)
		f /= math.Pow(10, float64(l-1))
		f = math.Floor(f*10.0+0.5) / 10.0

		idx := l/3 - 1
		if len(abbrSuffix) >= idx {
			return strconv.FormatFloat(f, 'f', -1, 64) + abbrSuffix[idx]
		}
	}

	return num
}
