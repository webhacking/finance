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

func ParseNumber(number string) (float64, error) {
	return strconv.ParseFloat(SanitizeNumber(number), 64)
}
