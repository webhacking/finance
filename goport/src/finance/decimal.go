package finance

// Decimal is a fixed-point type with four decimal points
type Decimal int64

const DecimalDigits = 4
const DecimalMultiplier = 10000

// DecimalFromString parses a string-represented number as Decimal
func DecimalFromString(value string) Decimal {
	parsed, _ := ParseNumber(value)
	// FIXME: Perhaps we should deal with the error here
	return Decimal(parsed * DecimalMultiplier)
}

func DecimalFromFloat(value float64) Decimal {
	return Decimal(value * DecimalMultiplier)
}

func (d Decimal) Floor() int64 {
	return int64(d) / DecimalMultiplier
}

// AsFloat converts a Decimal type to float64 (approximation)
// func (d Decimal) AsFloat() float64 {
// 	return float64(int64(d) / float64(DecimalMultiplier))
// }
