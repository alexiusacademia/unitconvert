package unitconvert

var Meter Unit = Unit{Name: "meter", Type: LengthUnit, System: "metric"}
var Feet Unit = Unit{Name: "feet", Type: LengthUnit, System: "english"}

// Length calculates the conversion of a value from one length unit to another.
// It takes a value (v) in the source unit (from) and converts it to the destination unit (to).
// Parameters:
//   - v: The value to be converted.
//   - from: The unit of the input value.
//   - to: The unit to which the input value is to be converted.
//
// Returns:
//   - float64: The converted value.
//   - error: An error indicating any conversion or input validation issues.
func Length(v float64, from, to Unit) (float64, error) {
	if from.Type != LengthUnit || to.Type != LengthUnit {
		return 0, invalidInputTypes
	}

	if from == to {
		return v, nil
	}

	factor, ok := conversionFactors[from][to]
	if !ok {
		return 0, unsupportedDestinationError
	}

	return v * factor, nil
}
