package finance

import (
	"strconv"
	"strings"
)

// SanitizeNumber clears commas in a decimal value in string type so that it
// can be parsed by strconv.ParseFloat().
func SanitizeNumber(number string) string {
	return strings.ReplaceAll(number, ",", "")
}

// ParseNumber parses a decimal value in string as a float64 type
func ParseNumber(number string) (float64, error) {
	return strconv.ParseFloat(SanitizeNumber(number), 64)
}
