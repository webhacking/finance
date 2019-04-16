package finance

import "testing"

func TestDecimalFromFloat(t *testing.T) {
	params := []struct {
		value    float64
		expected int64
	}{
		{0.0, 0},
		{0.2345, 2345},
		{1.3825, 13825},
		{17.0000, 170000},
		{845923.4952, 8459234952},
	}

	for _, param := range params {
		actual := DecimalFromFloat(param.value)
		expected := Decimal(param.expected)
		assertEquals(t, expected, actual, "Incorrect value")
	}
}
func TestDecimalFloor(t *testing.T) {
	params := []struct {
		value    int64
		expected int64
	}{
		{10000, 1},
		{12345, 1},
		{8459492, 845},
	}

	for _, param := range params {
		actual := Decimal(param.value).Floor()
		assertEquals(t, param.expected, actual, "Incorrect value")
	}
}
