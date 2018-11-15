package numba

import "strconv"

const (
	th = "th"
	st = "st"
	nd = "nd"
	rd = "rd"
)

// Ordinal formats integer n with its ordinal suffix.
func Ordinal(n int) string {
	return strconv.Itoa(n) + OrdinalSuffix(n)
}

// OrdinalSuffix returns the ordinal suffix that should be added to integer n.
func OrdinalSuffix(n int) string {
	switch n % 100 {
	case 11, 12, 13: // go switch cases do not fall through
	default:
		switch n % 10 {
		case 1:
			return st
		case 2:
			return nd
		case 3:
			return rd
		}
	}
	return th
}
