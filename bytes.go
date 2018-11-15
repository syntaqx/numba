package numba

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

type baseByteUnit = float64

// International Electrotechnical Commission (IEC) base 2 units [IEEE 1541]
// - http://physics.nist.gov/cuu/Units/binary.html
// - https://en.wikipedia.org/wiki/Binary_prefix
const (
	// Byte is the baseByteUnit for all binary prefix sizes
	Byte baseByteUnit = 1 << (iota * 10)
	// Kibibyte is a multiple of a Byte.
	Kibibyte
	// Mebibyte is a multiple of a Byte.
	Mebibyte
	// Gibibyte is a multiple of a Byte.
	Gibibyte
	// Tebibyte is a multiple of a Byte.
	Tebibyte
	// Pebibyte is a multiple of a Byte.
	Pebibyte
	// Exbibyte is a multiple of a Byte.
	Exbibyte
	// Zebibyte is a multiple of a Byte.
	Zebibyte
	// Yobibyte is a multiple of a Byte.
	Yobibyte
)

const (
	// KiB is the shorthand for Kibibyte
	KiB = Kibibyte
	// MiB is the shorthand for Mebibyte
	MiB = Mebibyte
	// GiB is the shorthand for Gibibyte
	GiB = Gibibyte
	// TiB is the shorthand for Tebibyte
	TiB = Tebibyte
	// PiB is the shorthand for Pebibyte
	PiB = Pebibyte
	// EiB is the shorthand for Exbibyte
	EiB = Exbibyte
	// ZiB is the shorthand for Zebibyte
	ZiB = Zebibyte
	// YiB is the shorthand for Yobibyte
	YiB = Yobibyte
)

// International Systems of Units (SI) base 10 units
// - http://physics.nist.gov/cuu/Units/prefixes.html
const (
	// Kilobyte is a multiple of a Byte.
	Kilobyte baseByteUnit = 1000 * Byte
	// Megabyte is a multiple of a Byte.
	Megabyte = 1000 * Kilobyte
	// Gigabyte is a multiple of a Byte.
	Gigabyte = 1000 * Megabyte
	// Terabyte is a multiple of a Byte.
	Terabyte = 1000 * Gigabyte
	// Petabyte is a multiple of a Byte.
	Petabyte = 1000 * Terabyte
	// Zettabyte is a multiple of a Byte.
	Exabyte = 1000 * Petabyte
	// Zettabyte is a multiple of a Byte.
	Zettabyte = 1000 * Exabyte
	// Yottabyte is a multiple of a Byte.
	Yottabyte = 1000 * Zettabyte
)

const (
	// KB is the shorthand for Kilobyte
	KB = Kilobyte
	// MB is the shorthand for Megabyte
	MB = Megabyte
	// GB is the shorthand for Gigabyte
	GB = Gigabyte
	// TB is the shorthand for Terabyte
	TB = Terabyte
	// PB is the shorthand for Petabyte
	PB = Petabyte
	// EB is the shorthand for Exabyte
	EB = Exabyte
	// ZB is the shorthand for Zettabyte
	ZB = Zettabyte
	// YB is the shorthand for Yottabyte
	YB = Yottabyte
)

var prefixes = map[string]float64{
	"KB":  KB,
	"MB":  MB,
	"GB":  GB,
	"TB":  TB,
	"PB":  PB,
	"EB":  EB,
	"ZB":  ZB,
	"YB":  YB,
	"KiB": KiB,
	"MiB": MiB,
	"GiB": GiB,
	"TiB": TiB,
	"PiB": PiB,
	"EiB": EiB,
	"ZiB": ZiB,
	"YiB": YiB,
}

// Bytes formats a given number of bytes in human-readable format, in either
// base 2 (IEC) or base 10 (SI) units as specified by the base parameter, with
// the specified number of digits of precision.
//
//   Bytes(1024*1024, 2, 2) => "1.00MiB"
//   Bytes(2000000000, 10, 2) => "2.00GB"
//
// If an invalid base or precision is given, a Sprintf-style error string is
// returned instead of the formatted bytes.
func Bytes(bytes int64, base int, prec int) string {
	if base != 10 && base != 2 {
		return "%!(BADBASE)"
	}

	b := float64(bytes)

	if base == 10 {
		switch {
		case b >= YB:
			return fmt.Sprintf("%.*fYB", prec, float64(b/YB))
		case b >= ZB:
			return fmt.Sprintf("%.*fZB", prec, float64(b/ZB))
		case b >= EB:
			return fmt.Sprintf("%.*fEB", prec, float64(b/EB))
		case b >= PB:
			return fmt.Sprintf("%.*fPB", prec, float64(b/PB))
		case b >= TB:
			return fmt.Sprintf("%.*fTB", prec, float64(b/TB))
		case b >= GB:
			return fmt.Sprintf("%.*fGB", prec, float64(b/GB))
		case b >= MB:
			return fmt.Sprintf("%.*fMB", prec, float64(b/MB))
		case b >= KB:
			return fmt.Sprintf("%.*fKB", prec, float64(b/KB))
		}
		return fmt.Sprintf("%.*fB", prec, float64(b))
	}

	switch {
	case b >= YiB:
		return fmt.Sprintf("%.*fYiB", prec, float64(b/YiB))
	case b >= ZiB:
		return fmt.Sprintf("%.*fZiB", prec, float64(b/ZiB))
	case b >= EiB:
		return fmt.Sprintf("%.*fEiB", prec, float64(b/EiB))
	case b >= PiB:
		return fmt.Sprintf("%.*fPiB", prec, float64(b/PiB))
	case b >= TiB:
		return fmt.Sprintf("%.*fTiB", prec, float64(b/TiB))
	case b >= GiB:
		return fmt.Sprintf("%.*fGiB", prec, float64(b/GiB))
	case b >= MiB:
		return fmt.Sprintf("%.*fMiB", prec, float64(b/MiB))
	case b >= KiB:
		return fmt.Sprintf("%.*fKiB", prec, float64(b/KiB))
	}

	return fmt.Sprintf("%.*fB", prec, float64(b))
}

func byteSplit(s string) (string, string) {
	s1 := make([]rune, 0, len(s))
	s2 := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) && len(s2) == 0 {
			s1 = append(s1, r)
		} else {
			if r != ' ' {
				s2 = append(s2, r)
			}
		}
	}
	return string(s1), string(s2)
}

// ParseBytesFloat parses a human-readable quantity of bytes, and returns the
// raw number of bytes as a float64.
//
// Whitespace is allowed between the number and the units. The 'B' for bytes is
// required to be upper case; 'b' in lower case would be bits.
func ParseBytesFloat(s string) (float64, error) {
	b, u := byteSplit(s)
	fb, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return fb, err
	}
	if u == "" {
		return fb, fmt.Errorf("no units found")
	}
	m, ok := prefixes[u]
	if !ok {
		return fb, fmt.Errorf("unrecognized units %s", u)
	}
	return m * fb, nil
}

// ParseBytes parses a human-readable quantity of bytes, and returns the raw
// number of bytes as an int64. If the value is too large for an int64, an error
// value is returned.
//
// Whitespace is allowed between the number and the units. The 'B' for bytes is
// required to be upper case; 'b' in lower case would be bits.
func ParseBytes(s string) (int64, error) {
	fv, err := ParseBytesFloat(s)
	if err != nil {
		return int64(fv), err
	}
	if fv > math.MaxInt64 {
		return 0, fmt.Errorf("value too large for int64")
	}
	return int64(fv), nil
}
