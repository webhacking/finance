package finance

import (
	"testing"
)

func TestParseNumber(t *testing.T) {
	var params = []struct {
		expected float64
		value    string
	}{
		{0, "0"},
		{123.45, "123.45"},
		{1000, "1,000"},
		{1234.567, "1,234.567"},
		{1000000, "1,000,000.00"},
	}

	for _, param := range params {
		actual, _ := ParseNumber(param.value)
		assertEquals(t, param.expected, actual, "Incorrect value")
	}
}
