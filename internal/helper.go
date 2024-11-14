package internal

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

func ConvertTaxStringToFloat64(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", ".")
	return strconv.ParseFloat(s, 64)
}

func ConvertStringToFloat64(s string) (float64, error) {
	// Remove currency symbol and any whitespace
	s = strings.TrimSpace(strings.TrimPrefix(s, "R$"))

	// Remove thousands separators
	s = strings.ReplaceAll(s, ".", "")

	// Replace comma with dot for decimal point
	s = strings.ReplaceAll(s, ",", ".")

	// Extract the numeric part
	re := regexp.MustCompile(`[\d.]+`)
	numStr := re.FindString(s)

	// Convert to float64
	value, err := strconv.ParseFloat(numStr, 64)

	if err != nil {
		return 0, err
	}

	// Handle suffix (M for million, B for billion, etc.)
	if strings.Contains(s, "K") {
		value *= 1e3
	} else if strings.Contains(s, "M") {
		value *= 1e6
	} else if strings.Contains(s, "B") {
		value *= 1e9
	}

	return value, nil
}

func Normalization(s string) string {
	t := norm.NFD.String(s)
	return strings.Map(func(r rune) rune {
		if unicode.IsMark(r) {
			return -1
		}
		return r
	}, t)
}
