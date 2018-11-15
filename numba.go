package numba

import (
	"math"
	"strconv"
)

type unit struct {
	Word string
	Abbr string
}

var units = [8]unit{
	{Word: "thousand", Abbr: "K"},
	{Word: "million", Abbr: "M"},
	{Word: "billion", Abbr: "B"},
	{Word: "trillion", Abbr: "T"},
	{Word: "quadrillion", Abbr: "P"},
	{Word: "quintillion", Abbr: "E"},
	{Word: "sextillion", Abbr: "Z"},
	{Word: "septillion", Abbr: "Y"},
}

func unitForInt(n int64) (string, *unit) {
	num := strconv.FormatInt(n, 10)
	l := len(num)

	if n > 0 && l > 3 {
		f := float64(n)
		f /= math.Pow(10, float64(l-1))
		f = math.Floor(f*10.0+0.5) / 10.0

		idx := l/3 - 1
		if len(units) >= idx {
			return strconv.FormatFloat(f, 'f', -1, 64), &units[idx]
		}
	}

	// fallback in the event we don't have a valid suffix
	return "", nil
}

// Word converts integer n to a friendly text word.
// For example, 1000 becomes "1 thousand" and 1500 becomes "1.5 thousand".
// If n is less than 100, the number is just returned as a string.
// Accepts octal or hex numbers, returned value is decimal in those cases too.
func Word(n int64) string {
	if v, u := unitForInt(n); u != nil {
		return v + " " + u.Word
	}
	return strconv.FormatInt(n, 10)
}

// NumAbbreviate converts integer n to a friendly word abbreviation.
// @see Word
func Abbreviate(n int64) string {
	if v, u := unitForInt(n); u != nil {
		return v + u.Abbr
	}
	return strconv.FormatInt(n, 10)
}
