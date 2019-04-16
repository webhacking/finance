package finance

import "math"

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
	return Decimal(math.Round(value * DecimalMultiplier))
}

func (d Decimal) Floor() int64 {
	return int64(d) / DecimalMultiplier
}

func (d Decimal) Ceil() int64 {
	value := int64(d)
	remainder := value % DecimalMultiplier

	if remainder > 0 {
		return value/DecimalMultiplier + 1
	} else {
		return value / DecimalMultiplier
	}
}

func (d Decimal) Round() int64 {
	value := int64(d)
	remainder := value % DecimalMultiplier

	if remainder > (DecimalMultiplier / 2) {
		return value/DecimalMultiplier + 1
	} else {
		return value / DecimalMultiplier
	}
}

// AsFloat converts a Decimal type to float64 (approximation)
func (d Decimal) AsFloat() float64 {
	return float64(d) / DecimalMultiplier
}

// TODO: We are going to need a DB adapter at some point in the future
